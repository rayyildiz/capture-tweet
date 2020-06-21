package main

import (
	"com.capturetweet/internal/convert"
	"com.capturetweet/pkg/service"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type handlerImpl struct {
	log     *zap.Logger
	service service.BrowserService
}

type response struct {
	Status  bool        `json:"status"`
	Code    *int        `json:"code,omitempty"`
	Message *string     `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (h handlerImpl) handleCapture(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&response{
			Status:  false,
			Code:    convert.Int(http.StatusMethodNotAllowed),
			Message: convert.String("method not allowed, only POST request"),
			Data:    nil,
		})
		h.log.Warn("method not allowed", zap.String("method", r.Method))

		return
	}

	request := service.CaptureRequestModel{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response{
			Status:  false,
			Code:    convert.Int(http.StatusBadRequest),
			Message: convert.String("bad request, check your input"),
			Data:    nil,
		})
		h.log.Error("bad request", zap.Error(err))
		return
	}
	defer r.Body.Close()

	err = h.service.CaptureSaveUpdateDatabase(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{
			Status:  false,
			Code:    convert.Int(http.StatusInternalServerError),
			Message: convert.String("could not capture your request, please try again"),
			Data:    nil,
		})
		h.log.Error("could not capture", zap.String("tweet_id", request.ID), zap.String("url", request.Url), zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&response{
		Status:  true,
		Code:    convert.Int(http.StatusOK),
		Message: nil,
		Data:    "capture saved in a bucket",
	})

}
