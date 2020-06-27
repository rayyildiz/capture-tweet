package main

import (
	"com.capturetweet/pkg/service"
	"encoding/json"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type handlerImpl struct {
	log     *zap.Logger
	service service.BrowserService
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
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		h.log.Error("bad request", zap.Error(err))
		return
	}
	defer r.Body.Close()

	request := service.CaptureRequestModel{}
	log.Printf("message: %v", payload)
	err = json.Unmarshal(payload.Message.Data, &request)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		h.log.Error("bad request, decode payload.data", zap.Error(err))
		return
	}

	respModel, err := h.service.CaptureSaveUpdateDatabase(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		h.log.Error("could not capture", zap.String("tweet_id", request.ID), zap.String("url", request.Url), zap.Error(err))
		return
	}

	h.log.Info("capture successfully", zap.String("tweet_id", respModel.ID), zap.String("tweet_url", request.Url),
		zap.String("capture_image", respModel.CaptureURL), zap.String("capture_thumb_image", respModel.CaptureThumbURL))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("No Content"))
}
