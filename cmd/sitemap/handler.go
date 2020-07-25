package main

import (
	"com.capturetweet/pkg/tweet"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"gocloud.dev/blob"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type handlerImpl struct {
	log    *zap.Logger
	repo   tweet.Repository
	bucket *blob.Bucket
}

func (h handlerImpl) handleRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		h.log.Warn("method not allowed", zap.String("method", r.Method))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	start := time.Now()

	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		h.log.Warn("there is no query string param for size", zap.String("path", r.URL.Path))
		size = 50
	}

	tweets, err := h.repo.FindAllOrderByUpdated(ctx, size)
	if err != nil {
		sentry.CaptureException(err)
		h.log.Warn("could not get repositories", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = createSitemap(ctx, h.log, h.bucket, tweets)
	if err != nil {
		sentry.CaptureException(err)
		h.log.Warn("could not get create sitemap", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	diff := time.Now().Sub(start)
	h.log.Info("create sitemap", zap.Duration("elapsed", diff))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("No Content"))
}

func createSitemap(ctx context.Context, log *zap.Logger, bucket *blob.Bucket, tweets []tweet.Tweet) error {
	sb := strings.Builder{}
	for _, tweet := range tweets {
		sb.WriteString(fmt.Sprintf(`
<url>
	<loc>https://capturetweet.com/tweet/%s</loc>
	<lastmod>%s</lastmod>
	<changefreq>weekly</changefreq>
	<priority>0.2</priority>
	<mobile:mobile/>
</url>`, tweet.ID, tweet.CreatedAt.Format(time.RFC3339)))
	}

	content := sb.String()

	xml := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
		xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd"
		xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
		xmlns:image="http://www.google.com/schemas/sitemap-image/1.1"
		xmlns:video="http://www.google.com/schemas/sitemap-video/1.1"
		xmlns:geo="http://www.google.com/geo/schemas/sitemap/1.0"
		xmlns:news="http://www.google.com/schemas/sitemap-news/0.9"
		xmlns:mobile="http://www.google.com/schemas/sitemap-mobile/1.0"
		xmlns:pagemap="http://www.google.com/schemas/sitemap-pagemap/1.0"
		xmlns:xhtml="http://www.w3.org/1999/xhtml"
	>
	<url>
		<loc>https://capturetweet.com/</loc>
		<lastmod>%s</lastmod>
		<changefreq>daily</changefreq>
		<priority>0.5</priority>
		<mobile:mobile/>
	</url>
	%s
</urlset>`, time.Now().Format(time.RFC3339), content)

	newSitemap := []byte(xml)

	oldSitemapAttrs, err := bucket.Attributes(ctx, "sitemap.xml")
	if err != nil {
		log.Error("bucket:ReadAll", zap.Error(err))
		return err
	}
	h := sha256.New()
	newHash := base64.StdEncoding.EncodeToString(h.Sum([]byte(content)))

	oldHash := oldSitemapAttrs.Metadata["x-content-sha256"]

	if newHash != oldHash {

		log.Info("old and new sitemap NOT equal, write on bucket")

		err = bucket.WriteAll(ctx, "sitemap.xml", newSitemap, &blob.WriterOptions{
			ContentType:  "application/xml",
			CacheControl: "public,max-age=9600",
			Metadata: map[string]string{
				"x-content-sha256": newHash,
			},
		})
		if err != nil {
			log.Error("bucket:writeAll", zap.Error(err))
			return err
		}

		log.Info("old and new sitemap NOT equal, ping search engines.")

		var wg sync.WaitGroup
		wg.Add(2)

		go func(ctx context.Context) {
			defer wg.Done()
			ping(ctx, "https://www.google.com/ping?sitemap=https://capturetweet.com/sitemap.xml", log)
		}(ctx)

		go func(ctx context.Context) {
			defer wg.Done()
			ping(ctx, "https://www.bing.com/ping?sitemap=https%3A%2F%2Fcapturetweet.com/sitemap.xml", log)
		}(ctx)
		wg.Wait()

	} else {
		log.Info("old and new sitemap equals, no need to ping search engines.", zap.String("old_hash", oldHash), zap.String("new_hash", newHash))
	}
	return nil
}

func ping(ctx context.Context, pingUrl string, log *zap.Logger) {
	request, err := http.NewRequest(http.MethodGet, pingUrl, nil)
	if err != nil {
		log.Error("error creating request", zap.String("url", pingUrl), zap.Error(err))
		return
	}
	request.Header.Add("x-app-name", "go-http")
	request.Header.Add("x-app-version", "0.3.0")
	request = request.WithContext(ctx)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Error("error while ping", zap.String("url", pingUrl), zap.Error(err))
	} else {
		log.Info("success", zap.String("url", pingUrl), zap.String("status", resp.Status))
	}
}
