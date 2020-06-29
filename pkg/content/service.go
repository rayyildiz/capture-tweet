package content

import (
	"bytes"
	"com.capturetweet/pkg/service"
	"fmt"
	"log"
	"net/http"
	"os"
)

type serviceImpl struct {
	apiKey string
}

func NewService(apiKey string) service.ContentService {
	return &serviceImpl{apiKey}
}

func (s serviceImpl) SendMail(senderMail, senderName, message string) error {
	body := bytes.NewBufferString(fmt.Sprintf(`{
  "personalizations": [
    {
      "to": [
				{
          "email": "rayyildiz@ymail.com",
          "name": "Ramazan AYYILDIZ"
        }
      ],
      "dynamic_template_data": {
        "name": "%s",
        "mail": "%s",
        "message": "%s"
      },
      "subject": "[CaptureTweet] Contact us"
    }
  ],
  "from": {
    "email": "%s",
    "name": "%s"
  },
  "template_id": "%s"
}`, senderName, senderMail, message, senderMail, senderName, os.Getenv("SENDGRID_TEMPLATE_ID")))

	request, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", body)
	if err != nil {
		return err
	}
	request.Header.Set("authorization", "Bearer "+s.apiKey)
	request.Header.Set("content-type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	log.Printf("http send response %v", response.Status)
	return nil
}
