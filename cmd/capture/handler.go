package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"capturetweet.com/api"
)

type handlerImpl struct {
	service api.BrowserService
}

type PubSubMessage struct {
	Message struct {
		Data       []byte            `json:"data"`
		MessageId  string            `json:"messageId"`
		Attributes map[string]string `json:"attributes"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func (h handlerImpl) handleCapture(w http.ResponseWriter, r *http.Request) {
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

	request := api.CaptureRequestModel{}
	err = json.Unmarshal(payload.Message.Data, &request)
	if err != nil {
		slog.Error("bad request, decode payload.data", slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	respModel, err := h.service.CaptureSaveUpdateDatabase(r.Context(), &request)
	if err != nil {
		slog.Error("could not capture", slog.String("tweet_id", request.ID), slog.String("url", request.Url), slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	slog.Info("capture successfully", slog.String("tweet_id", respModel.ID), slog.String("tweet_url", request.Url), slog.String("capture_image", respModel.CaptureURL))
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("No Content"))
}
