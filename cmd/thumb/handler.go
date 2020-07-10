package main

import (
	"bytes"
	"com.capturetweet/api"
	"context"
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/nfnt/resize"
	"go.uber.org/zap"
	"gocloud.dev/blob"
	"image"
	"image/jpeg"
	"net/http"
	"strings"
	"time"
)

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

type StorageMessage struct {
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	ContentType string `json:"contentType"`
	Size        string `json:"size"`
}

type handlerImpl struct {
	log     *zap.Logger
	service api.TweetService
	bucket  *blob.Bucket
}

func (h handlerImpl) handleResize(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		h.log.Warn("method not allowed", zap.String("method", r.Method))
		return
	}

	var payload PubSubMessage
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		sentry.CaptureException(err)
		h.log.Error("bad request", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	request := StorageMessage{}
	err = json.Unmarshal(payload.Message.Data, &request)
	if err != nil {
		sentry.CaptureException(err)
		h.log.Error("bad request, decode payload.data", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Kind != "storage#object" {
		h.log.Warn("expected image kind must be object", zap.String("image_kind", request.Kind), zap.String("image_key", request.Name))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("it is not an object"))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	img, err := h.bucket.ReadAll(ctx, request.Name)
	if err != nil {
		sentry.CaptureException(err)
		h.log.Error("open bucket", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	attrs, err := h.bucket.Attributes(ctx, request.Name)
	if err != nil {
		sentry.CaptureException(err)
		h.log.Error("read image attributes", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	tweetId := attrs.Metadata["tweet_id"]
	tweetUser := attrs.Metadata["tweet_user"]
	tweetUrl := attrs.Metadata["tweet_url"]

	decoder, _, err := image.Decode(bytes.NewBuffer(img))
	if err != nil {
		sentry.CaptureException(err)
		h.log.Error("image decode", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	split := strings.Split(request.Name, "/")
	fileName := split[len(split)-1]
	thumbNailKey := fmt.Sprintf("capture/thumb/%s", fileName)

	newImage := resize.Resize(320, 0, decoder, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, newImage, nil)
	if err != nil {
		sentry.CaptureException(err)
		h.log.Error("jpeg encode image", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
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
		sentry.CaptureException(err)
		h.log.Error("open bucket for thumbnail", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	h.log.Info("image stored successfully", zap.String("image_key", thumbNailKey), zap.String("image_key", request.Name), zap.String("tweet_id", tweetId), zap.String("tweet_user", tweetUser))
	err = h.service.UpdateThumbImage(tweetId, thumbNailKey)
	if err != nil {
		sentry.CaptureException(err)
		h.log.Error("save in database", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	diff := time.Now().Sub(start)
	h.log.Info("image saved", zap.Duration("elapsed", diff), zap.String("image_thumb", thumbNailKey), zap.String("image_key", request.Name), zap.String("image_kind", request.Kind))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("No Content"))
}
