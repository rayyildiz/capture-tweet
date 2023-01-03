package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"strings"
	"time"

	"capturetweet.com/api"
	"github.com/nfnt/resize"
	"go.uber.org/zap"
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
	zap.L().Info(s)
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		zap.L().Error("can not read ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	zap.L().Info("aaa", zap.ByteString("data", b))
}

func (h handlerImpl) handleResize(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		zap.L().Warn("method not allowed", zap.String("method", r.Method))
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var payload EventArcRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		zap.L().Error("bad request", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if !strings.HasPrefix(payload.Name, "capture/large/") {
		zap.L().Warn("not a valid prefix", zap.String("image_kind", payload.Kind), zap.String("name", payload.Name), zap.String("id", payload.Id))
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(w, "")
		return
	}

	if payload.Kind != "storage#object" {
		zap.L().Warn("expected image kind must be object", zap.String("image_kind", payload.Kind), zap.String("name", payload.Name), zap.String("id", payload.Id))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("it is not an object"))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	img, err := h.bucket.ReadAll(ctx, payload.Name)
	if err != nil {
		zap.L().Error("open bucket", zap.String("image_key", payload.Name), zap.String("image_kind", payload.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	attrs, err := h.bucket.Attributes(ctx, payload.Name)
	if err != nil {
		zap.L().Error("read image attributes", zap.String("image_key", payload.Name), zap.String("image_kind", payload.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	tweetId := attrs.Metadata["tweet_id"]
	tweetUser := attrs.Metadata["tweet_user"]
	tweetUrl := attrs.Metadata["tweet_url"]

	decoder, _, err := image.Decode(bytes.NewBuffer(img))
	if err != nil {
		zap.L().Error("image decode", zap.String("image_key", payload.Name), zap.String("image_kind", payload.Kind), zap.Error(err))
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
		zap.L().Error("jpeg encode image", zap.String("image_key", payload.Name), zap.String("image_kind", payload.Kind), zap.Error(err))
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
		zap.L().Error("open bucket for thumbnail", zap.String("image_key", payload.Name), zap.String("image_kind", payload.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	zap.L().Info("image stored successfully", zap.String("image_key", thumbNailKey), zap.String("image_key", payload.Name), zap.String("tweet_id", tweetId), zap.String("tweet_user", tweetUser))
	err = h.service.UpdateThumbImage(ctx, tweetId, thumbNailKey)
	if err != nil {
		zap.L().Error("save in database", zap.String("image_key", payload.Name), zap.String("image_kind", payload.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	zap.L().Info("image saved", zap.Duration("elapsed", time.Since(start).Round(time.Millisecond)), zap.String("image_thumb", thumbNailKey), zap.String("image_key", payload.Name), zap.String("image_kind", payload.Kind))

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("No Content"))
}
