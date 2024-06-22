package main

import (
	"encoding/json"
	"fmt"
	"gocloud.dev/blob"
	"log/slog"
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
		slog.Warn("method not allowed", slog.String("method", r.Method))
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var payload PubSubMessage
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		slog.Error("bad request", slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	key := fmt.Sprintf("dlq/%s.json", payload.Message.MessageId)
	slog.Info("message saving into bucket", slog.String("key", key))

	err = h.bucket.WriteAll(r.Context(), key, payload.Message.Data, &blob.WriterOptions{
		ContentType: "application/json",
		Metadata: map[string]string{
			"message-id":   payload.Message.MessageId,
			"subscription": payload.Subscription,
			"publish-time": payload.PublishTime,
		},
	})
	if err != nil {
		slog.Error("could not create a write", slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
