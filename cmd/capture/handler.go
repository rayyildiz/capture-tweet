package main

import (
	"encoding/json"
	"net/http"

	"capturetweet.com/api"
	"go.uber.org/zap"
)

type handlerImpl struct {
	service api.BrowserService
}

type PubSubMessage struct {
	Subscription string `json:"subscription"`
	Message      struct {
		Attributes map[string]string `json:"attributes"`
		MessageId  string            `json:"messageId"`
		Data       []byte            `json:"data"`
	} `json:"message"`
}

func (h handlerImpl) handleCapture(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		zap.L().Warn("method not allowed", zap.String("method", r.Method))
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
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

	request := api.CaptureRequestModel{}
	err = json.Unmarshal(payload.Message.Data, &request)
	if err != nil {
		zap.L().Error("bad request, decode payload.data", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	respModel, err := h.service.CaptureSaveUpdateDatabase(r.Context(), &request)
	if err != nil {
		zap.L().Error("could not capture", zap.String("tweet_id", request.ID), zap.String("url", request.Url), zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	zap.L().Info("capture successfully", zap.String("tweet_id", respModel.ID), zap.String("tweet_url", request.Url), zap.String("capture_image", respModel.CaptureURL))
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("No Content"))
}
