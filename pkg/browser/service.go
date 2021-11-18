package browser

import (
	"context"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"capturetweet.com/api"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"go.uber.org/zap"
	"gocloud.dev/blob"
)

var minImageSize = 1024 * 25

type serviceImpl struct {
	tweetService api.TweetService
	bucket       *blob.Bucket
	browser      *browserCtx
}

type browserCtx struct {
	browserContext context.Context
	cancelFunc     context.CancelFunc
}

func NewService(tweetService api.TweetService, bucket *blob.Bucket) api.BrowserService {
	return &serviceImpl{
		tweetService: tweetService,
		bucket:       bucket,
		browser:      nil,
	}
}

func (s serviceImpl) CaptureSaveUpdateDatabase(ctx context.Context, model *api.CaptureRequestModel) (*api.CaptureResponseModel, error) {
	originalImage, err := s.CaptureURL(ctx, model)
	if err != nil {
		zap.L().Error("browser CaptureSaveUpdateDatabase, captureURL", zap.String("tweet_id", model.ID), zap.Error(err))
		return nil, err
	}

	if len(originalImage) < minImageSize {
		return nil, fmt.Errorf("image size is less than 25 Kb. try again. size is %d", len(originalImage))
	}

	response, err := s.SaveCapture(ctx, originalImage, model)
	if err != nil {
		zap.L().Error("browser CaptureSaveUpdateDatabase, SaveCapture", zap.String("tweet_id", model.ID), zap.Error(err))
		return nil, err
	}

	err = s.tweetService.UpdateLargeImage(ctx, model.ID, response.CaptureURL)
	if err != nil {
		zap.L().Error("browser CaptureSaveUpdateDatabase, service.UpdateLargeImage", zap.String("tweet_id", model.ID), zap.Error(err))
		return nil, err
	}

	return response, nil
}

func (s serviceImpl) SaveCapture(ctx context.Context, originalImage []byte, model *api.CaptureRequestModel) (*api.CaptureResponseModel, error) {
	imageKey := fmt.Sprintf("capture/large/%s.jpg", model.ID)

	err := s.bucket.WriteAll(ctx, imageKey, originalImage, &blob.WriterOptions{
		ContentType:  "image/jpg",
		CacheControl: "private,max-age=86400",
		Metadata: map[string]string{
			"image_type": "original",
			"tweet_id":   model.ID,
			"tweet_url":  model.Url,
			"tweet_user": model.Author,
		},
	})
	if err != nil {
		zap.L().Error("browser:saveCapture", zap.String("tweet_id", model.ID), zap.Error(err))
		return nil, err
	}

	return &api.CaptureResponseModel{
		ID:         model.ID,
		CaptureURL: imageKey,
	}, nil
}
func (s serviceImpl) Close() {
	if s.browser != nil {
		zap.L().Info("closing browser context")
		s.browser.cancelFunc()
	}
}

func (s *serviceImpl) CaptureURL(ctx context.Context, model *api.CaptureRequestModel) ([]byte, error) {
	if s.browser == nil {
		opts := []chromedp.ExecAllocatorOption{
			chromedp.DisableGPU,
			chromedp.Headless,
			chromedp.NoDefaultBrowserCheck,
			chromedp.NoSandbox,
			chromedp.Flag("disable-translate", true),
			chromedp.Flag("disable-sync", true),
			chromedp.Flag("user-data-dir", "/tmp/headless"),
		}

		allocCtx, _ := chromedp.NewExecAllocator(ctx, opts...)
		ctx, cancel := chromedp.NewContext(allocCtx)
		s.browser = &browserCtx{
			browserContext: ctx,
			cancelFunc:     cancel,
		}
	}

	var buf []byte
	err := chromedp.Run(s.browser.browserContext, fullScreenshot(model.Url, 90, &buf))
	if err != nil {
		zap.L().Error("could not capture URL", zap.String("tweet_id", model.ID), zap.String("url", model.Url), zap.Error(err))
		return nil, err
	}
	return buf, nil
}

func getSleep() time.Duration {
	ms, _ := strconv.Atoi(os.Getenv("APP_SLEEP_TIME_MS"))
	if ms == 0 {
		ms = 3000
	}
	return time.Duration(ms) * time.Millisecond
}

func fullScreenshot(url string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(getSleep()),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, _, _, _, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}
			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithFormat(page.CaptureScreenshotFormatJpeg).
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
