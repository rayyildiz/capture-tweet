//go:build wireinject
// +build wireinject

package main

import (
	"capturetweet.com/api"
	"capturetweet.com/internal/infra"
	"capturetweet.com/pkg/browser"
	"capturetweet.com/pkg/tweet"
	"github.com/google/wire"
)

func initializeBrowserService() api.BrowserService {
	wire.Build(infra.NewPostgresDatabase, infra.NewBucketFromEnvironment, tweet.NewRepository, tweet.NewServiceWithRepository, browser.NewService)

	return nil
}
