package browser

import (
	"context"
	"os"
	"testing"

	"capturetweet.com/api"
	"capturetweet.com/internal/infra"
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

	infra.RegisterLogger()

	service := NewService(nil, nil)
	require.NotNil(t, service)

	minImageSize = 1024 * 10

	data, err := service.CaptureURL(ctx, &api.CaptureRequestModel{
		ID:     "1335519471822381056",
		Author: "rayyildiz",
		Url:    "https://twitter.com/rayyildiz/status/1335519471822381056",
	})
	require.NoError(t, err)
	if assert.NotNil(t, data) {
		assert.True(t, len(data) > 1000)
	}
}

func TestService_SaveCapture(t *testing.T) {
	ctx := context.Background()

	infra.RegisterLogger()

	bucket := infra.NewBucket("mem://test")
	require.NotNil(t, bucket)

	service := NewService(nil, bucket)
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
	os.Setenv("APP_SLEEP_TIME_MS", "9000")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "chromedp/headless-shell",
		ExposedPorts: []string{"9222/tcp"},
		WaitingFor:   wait.ForListeningPort("9222"),
	}
	chromeC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	defer chromeC.Terminate(ctx)

	bucket := infra.NewBucket("mem://mem")
	require.NotNil(t, bucket)

	infra.RegisterLogger()

	tweetS := api.NewMockTweetService(ctrl)
	tweetS.EXPECT().UpdateLargeImage(gomock.Any(), "1356685552276434946", "capture/large/1356685552276434946.jpg").Return(nil)

	service := NewService(tweetS, bucket)
	require.NotNil(t, service)

	response, err := service.CaptureSaveUpdateDatabase(ctx, &api.CaptureRequestModel{
		ID:     "1356685552276434946",
		Author: "CloudNativeFdn",
		Url:    "https://twitter.com/CloudNativeFdn/status/1356685552276434946",
	})
	require.NoError(t, err)
	if assert.NotNil(t, response) {
	}
}
