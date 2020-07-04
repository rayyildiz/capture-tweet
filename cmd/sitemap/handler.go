package main

import (
	"bytes"
	"com.capturetweet/pkg/tweet"
	"context"
	"fmt"
	"go.uber.org/zap"
	"gocloud.dev/blob"
	"net/http"
	"strings"
	"time"
)

type handlerImpl struct {
	log    *zap.Logger
	repo   tweet.Repository
	bucket *blob.Bucket
}

func (h handlerImpl) handleRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		h.log.Warn("method not allowed", zap.String("method", r.Method))
		return
	}

	start := time.Now()

	tweets, err := h.repo.FindAllOrderByUpdated(50)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		h.log.Warn("could not get repositories", zap.Error(err))
		return
	}

	err = createSitemap(r.Context(), h.log, h.bucket, tweets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		h.log.Warn("could not get create sitemap", zap.Error(err))
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
</urlset>`, time.Now().Format(time.RFC3339), sb.String())

	newSitemap := []byte(xml)

	oldSitemap, err := bucket.ReadAll(ctx, "sitemap.xml")
	if err != nil {
		log.Error("bucket:ReadAll", zap.Error(err))
		return err
	}

	err = bucket.WriteAll(ctx, "sitemap.xml", newSitemap, &blob.WriterOptions{
		ContentType:  "application/xml",
		CacheControl: "public,max-age=86400",
	})
	if err != nil {
		log.Error("bucket:writeAll", zap.Error(err))
		return err
	}

	if bytes.Compare(newSitemap, oldSitemap) != 0 {
		log.Info("old and new sitemap NOT equal, ping search engines.")
		go ping("https://www.google.com/ping?sitemap=https://capturetweet.com/sitemap.xml", log)
		go ping("https://www.bing.com/ping?sitemap=https%3A%2F%2Fcapturetweet.com/sitemap.xml", log)
	} else {
		log.Info("old and new sitemap equals, no need to ping search engines.")
	}
	return nil
}

func ping(pingUrl string, log *zap.Logger) {
	resp, err := http.Get(pingUrl)
	if err != nil {
		log.Error("error while ping", zap.String("url", pingUrl), zap.Error(err))
	} else {
		log.Info("success", zap.String("url", pingUrl), zap.String("status", resp.Status))
	}
}
