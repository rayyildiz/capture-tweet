//go:build wireinject
// +build wireinject

package main

import (
	"com.capturetweet/api"
	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/browser"
	"com.capturetweet/pkg/tweet"
	"github.com/google/wire"
)

func initializeBrowserService() api.BrowserService {
	wire.Build(infra.NewTweetCollection, infra.NewBucketFromEnvironment, tweet.NewRepository, tweet.NewServiceWithRepository, browser.NewService)

	return nil
}
