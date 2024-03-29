package content

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"capturetweet.com/api"
)

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) api.ContentService {
	return &serviceImpl{
		repo: repo,
	}
}

type captchaResponse struct {
	Success    bool     `json:"success"`
	ErrorCodes []string `json:"error-codes"`
}

func (s serviceImpl) StoreContactRequest(ctx context.Context, senderMail, senderName, message, captcha string) error {
	post := url.Values{
		"secret":   {os.Getenv("CAPTCHA_SECRET")},
		"response": {captcha},
	}

	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", post)
	if err != nil {
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
