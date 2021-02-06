package main

import (
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"gocloud.dev/blob"
	"net/http"
)

type handlerImpl struct {
	bucket *blob.Bucket
}

type PubSubMessage struct {
	Message struct {
		Data       []byte            `json:"data"`
		MessageId  string            `json:"messageId"`
		Attributes map[string]string `json:"attributes"`
	} `json:"message"`
	Subscription string `json:"subscription"`
	PublishTime  string `json:"publishTime"`
}

func (h handlerImpl) handleMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		zap.L().Warn("method not allowed", zap.String("method", r.Method))
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var payload PubSubMessage
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("bad request", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	key := fmt.Sprintf("dlq/%s.json", payload.Message.MessageId)
	zap.L().Info("message saving into bucket", zap.String("key", key))

	err = h.bucket.WriteAll(r.Context(), key, payload.Message.Data, &blob.WriterOptions{
		ContentType: "application/json",
		Metadata: map[string]string{
			"message-id":   payload.Message.MessageId,
			"subscription": payload.Subscription,
			"publish-time": payload.PublishTime,
		},
	})
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("could not create a write", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
