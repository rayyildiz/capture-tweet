package main

import (
	"bytes"
	"com.capturetweet/api"
	"context"
	"encoding/json"
	"fmt"
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
	service api.TweetService
	bucket  *blob.Bucket
}

func (h handlerImpl) handleResize(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		zap.L().Warn("method not allowed", zap.String("method", r.Method))
		return
	}

	var payload PubSubMessage
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		zap.L().Error("bad request", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	request := StorageMessage{}
	err = json.Unmarshal(payload.Message.Data, &request)
	if err != nil {
		zap.L().Error("bad request, decode payload.data", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Kind != "storage#object" {
		zap.L().Warn("expected image kind must be object", zap.String("image_kind", request.Kind), zap.String("image_key", request.Name))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("it is not an object"))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	img, err := h.bucket.ReadAll(ctx, request.Name)
	if err != nil {
		zap.L().Error("open bucket", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	attrs, err := h.bucket.Attributes(ctx, request.Name)
	if err != nil {
		zap.L().Error("read image attributes", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	tweetId := attrs.Metadata["tweet_id"]
	tweetUser := attrs.Metadata["tweet_user"]
	tweetUrl := attrs.Metadata["tweet_url"]

	decoder, _, err := image.Decode(bytes.NewBuffer(img))
	if err != nil {
		zap.L().Error("image decode", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
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
		zap.L().Error("jpeg encode image", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
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
		zap.L().Error("open bucket for thumbnail", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	zap.L().Info("image stored successfully", zap.String("image_key", thumbNailKey), zap.String("image_key", request.Name), zap.String("tweet_id", tweetId), zap.String("tweet_user", tweetUser))
	err = h.service.UpdateThumbImage(ctx, tweetId, thumbNailKey)
	if err != nil {
		zap.L().Error("save in database", zap.String("image_key", request.Name), zap.String("image_kind", request.Kind), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	diff := time.Now().Sub(start)
	zap.L().Info("image saved", zap.Duration("elapsed", diff), zap.String("image_thumb", thumbNailKey), zap.String("image_key", request.Name), zap.String("image_kind", request.Kind))

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("No Content"))
}
