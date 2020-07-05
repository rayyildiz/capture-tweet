package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Tweet struct {
	User string
	Id   string
}

func main() {

	tweets, err := readFile()
	if err != nil {
		log.Fatalf("could not open file %v", err)
	}

	log.Printf("number of message %d", len(tweets))

	for _, tweet := range tweets {
		user, id := tweet.User, tweet.Id

		err := doRequest(user, id)
		if err != nil {
			log.Printf("[ERROR] %s/%s error, %v", user, id, err)
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func readFile() ([]Tweet, error) {
	file, err := ioutil.ReadFile("cmd/migration/results-20200704-122427.json")
	if err != nil {
		return nil, err
	}

	var tweets []Tweet

	err = json.Unmarshal(file, &tweets)

	return tweets, err
}

func doRequest(user, id string) error {
	body := fmt.Sprintf(`{"operationName":"Capture","variables":{"url":"https://twitter.com/%s/status/%s"},"query":"mutation Capture($url: String!) {  capture(url: $url) { id fullText favoriteCount retweetCount postedAt __typename } }"}`, user, id)

	req, err := http.NewRequest(http.MethodPost, "https://beta-api.capturetweet.com/api/query", bytes.NewBufferString(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user-agent", "go-migration")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO]  %s/%s reuest => %s", user, id, response.Status)

	return nil
}
