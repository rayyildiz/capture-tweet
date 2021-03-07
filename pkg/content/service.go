package content

import (
	"com.capturetweet/api"
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"net/url"
	"os"
)

type serviceImpl struct {
	repo   Repository
	tracer trace.Tracer
}

func NewService(repo Repository) api.ContentService {
	return &serviceImpl{
		repo:   repo,
		tracer: otel.GetTracerProvider().Tracer("com.capturetweet/pkg/content"),
	}
}

type captchaResponse struct {
	Success    bool     `json:"success"`
	ErrorCodes []string `json:"error-codes"`
}

func (s serviceImpl) StoreContactRequest(ctx context.Context, senderMail, senderName, message, captcha string) error {
	ctx, span := s.tracer.Start(ctx, "service:storeContactRequest")
	defer span.End()

	post := url.Values{
		"secret":   {os.Getenv("CAPTCHA_SECRET")},
		"response": {captcha},
	}

	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", post)
	if err != nil {
		span.RecordError(err)
		return err
	}
	defer resp.Body.Close()

	var r captchaResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return err
	}

	if !r.Success {
		return fmt.Errorf("captch failed, %v", r.ErrorCodes)
	}

	return s.repo.ContactUs(ctx, senderMail, senderName, message)
}
