package browser

import (
	"context"
	"os"
	"testing"

	"com.capturetweet/api"
	"com.capturetweet/internal/infra"
	"github.com/docker/go-connections/nat"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestService_CaptureURL(t *testing.T) {
	os.Setenv("APP_SLEEP_TIME_MS", "6000")
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "chromedp/headless-shell",
		ExposedPorts: []string{"9222/tcp"},
		WaitingFor:   wait.ForListeningPort(nat.Port("9222")),
	}
	chromeC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	defer chromeC.Terminate(ctx)

	log := infra.NewLogger()
	require.NotNil(t, log)

	service := NewService(log, nil, nil)
	require.NotNil(t, service)

	minImageSize = 1024 * 10

	data, err := service.CaptureURL(ctx, &api.CaptureRequestModel{
		ID:     "20",
		Author: "jack",
		Url:    "https://twitter.com/jack/20",
	})
	require.NoError(t, err)
	if assert.NotNil(t, data) {
		assert.True(t, len(data) > 1000)
	}
}

func TestService_SaveCapture(t *testing.T) {
	ctx := context.Background()

	log := infra.NewLogger()
	require.NotNil(t, log)

	bucket, err := infra.NewBucket("mem://test")
	require.NoError(t, err)
	require.NotNil(t, bucket)

	service := NewService(log, nil, bucket)
	require.NotNil(t, service)

	response, err := service.SaveCapture(ctx, []byte("hello"), &api.CaptureRequestModel{
		ID:     "1",
		Author: "example",
		Url:    "https://twitter.com/example/1",
	})
	require.NoError(t, err)
	if assert.NotNil(t, response) {
		assert.Equal(t, "1", response.ID)
		assert.Equal(t, "capture/large/1.jpg", response.CaptureURL)
	}
}

func TestService_CaptureSaveUpdateDatabase(t *testing.T) {
	os.Setenv("APP_SLEEP_TIME_MS", "6000")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "chromedp/headless-shell",
		ExposedPorts: []string{"9222/tcp"},
		WaitingFor:   wait.ForListeningPort(nat.Port("9222")),
	}
	chromeC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	defer chromeC.Terminate(ctx)

	bucket, err := infra.NewBucket("mem://mem2")
	require.NoError(t, err)
	require.NotNil(t, bucket)

	log := infra.NewLogger()
	require.NotNil(t, log)

	tweetS := api.NewMockTweetService(ctrl)
	tweetS.EXPECT().UpdateLargeImage(ctx, "20", "capture/large/20.jpg").Return(nil)

	service := NewService(log, tweetS, bucket)
	require.NotNil(t, service)

	response, err := service.CaptureSaveUpdateDatabase(ctx, &api.CaptureRequestModel{
		ID:     "20",
		Author: "jack",
		Url:    "https://twitter.com/jack/20",
	})
	require.NoError(t, err)
	if assert.NotNil(t, response) {
	}
}
