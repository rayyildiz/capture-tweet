package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"capturetweet.com/api"
	"github.com/nfnt/resize"
	"gocloud.dev/blob"
)

type EventArcRequest struct {
	Kind                    string    `json:"kind"`
	Id                      string    `json:"id"`
	SelfLink                string    `json:"selfLink"`
	Name                    string    `json:"name"`
	Bucket                  string    `json:"bucket"`
	Generation              string    `json:"generation"`
	Metageneration          string    `json:"metageneration"`
	ContentType             string    `json:"contentType"`
	TimeCreated             time.Time `json:"timeCreated"`
	Updated                 time.Time `json:"updated"`
	StorageClass            string    `json:"storageClass"`
	TimeStorageClassUpdated time.Time `json:"timeStorageClassUpdated"`
	Size                    string    `json:"size"`
	Md5Hash                 string    `json:"md5Hash"`
	MediaLink               string    `json:"mediaLink"`
	Crc32C                  string    `json:"crc32c"`
	Etag                    string    `json:"etag"`
}

type handlerImpl struct {
	service api.TweetService
	bucket  *blob.Bucket
}

func (h handlerImpl) handleLog(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Detected change in Cloud Storage bucket: %s", r.Header.Get("Ce-Subject"))
	slog.Info(s)
	defer r.Body.Close()

	_, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("can not read ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h handlerImpl) handleResize(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		slog.Warn("method not allowed", slog.String("method", r.Method))
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var payload EventArcRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		slog.Error("bad request", slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if !strings.HasPrefix(payload.Name, "capture/large/") {
		slog.Warn("not a valid prefix", slog.String("image_kind", payload.Kind), slog.String("name", payload.Name), slog.String("id", payload.Id))
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(w, "")
		return
	}

	if payload.Kind != "storage#object" {
		slog.Warn("expected image kind must be object", slog.String("image_kind", payload.Kind), slog.String("name", payload.Name), slog.String("id", payload.Id))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("it is not an object"))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	img, err := h.bucket.ReadAll(ctx, payload.Name)
	if err != nil {
		slog.Error("open bucket", slog.String("image_key", payload.Name), slog.String("image_kind", payload.Kind), slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	attrs, err := h.bucket.Attributes(ctx, payload.Name)
	if err != nil {
		slog.Error("read image attributes", slog.String("image_key", payload.Name), slog.String("image_kind", payload.Kind), slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	tweetId := attrs.Metadata["tweet_id"]
	tweetUser := attrs.Metadata["tweet_user"]
	tweetUrl := attrs.Metadata["tweet_url"]

	decoder, _, err := image.Decode(bytes.NewBuffer(img))
	if err != nil {
		slog.Error("image decode", slog.String("image_key", payload.Name), slog.String("image_kind", payload.Kind), slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	split := strings.Split(payload.Name, "/")
	fileName := split[len(split)-1]
	thumbNailKey := fmt.Sprintf("capture/thumb/%s", fileName)

	newImage := resize.Resize(320, 0, decoder, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, newImage, nil)
	if err != nil {
		slog.Error("jpeg encode image", slog.String("image_key", payload.Name), slog.String("image_kind", payload.Kind), slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = h.bucket.WriteAll(ctx, thumbNailKey, buf.Bytes(), &blob.WriterOptions{
		ContentType:  "image/jpg",
		CacheControl: "private,sitemap86400",
		Metadata: map[string]string{
			"image_type": "thumb",
			"tweet_id":   tweetId,
			"tweet_url":  tweetUrl,
			"tweet_user": tweetUser,
		},
	})
	if err != nil {
		slog.Error("open bucket for thumbnail", slog.String("image_key", payload.Name), slog.String("image_kind", payload.Kind), slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	slog.Info("image stored successfully", slog.String("image_key", thumbNailKey), slog.String("image_key", payload.Name), slog.String("tweet_id", tweetId), slog.String("tweet_user", tweetUser))
	err = h.service.UpdateThumbImage(ctx, tweetId, thumbNailKey)
	if err != nil {
		slog.Error("save in database", slog.String("image_key", payload.Name), slog.String("image_kind", payload.Kind), slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	slog.Info("image saved", slog.Duration("elapsed", time.Since(start).Round(time.Millisecond)), slog.String("image_thumb", thumbNailKey), slog.String("image_key", payload.Name), slog.String("image_kind", payload.Kind))

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("No Content"))
}
