package browser

import (
	"com.capturetweet/pkg/service"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"go.uber.org/zap"
	"gocloud.dev/blob"
	"math"
	"time"
)

type serviceImpl struct {
	log          *zap.Logger
	tweetService service.TweetService
	bucket       *blob.Bucket
	browser      *browserCtx
}

type browserCtx struct {
	browserContext context.Context
	cancelFunc     context.CancelFunc
}

func NewService(log *zap.Logger, tweetService service.TweetService, bucket *blob.Bucket) service.BrowserService {
	return &serviceImpl{log, tweetService, bucket, nil}
}

func (s serviceImpl) CaptureSaveUpdateDatabase(model *service.CaptureRequestModel) (*service.CaptureResponseModel, error) {

	originalImage, err := s.CaptureURL(model)
	if err != nil {
		s.log.Error("browser CaptureSaveUpdateDatabase, captureURL", zap.String("tweet_id", model.ID), zap.Error(err))
		return nil, err
	}

	if len(originalImage) < 1024*25 {
		return nil, fmt.Errorf("image size is less than 25 Kb. try again. size is %d", len(originalImage))
	}

	response, err := s.SaveCapture(originalImage, model)
	if err != nil {
		s.log.Error("browser CaptureSaveUpdateDatabase, SaveCapture", zap.String("tweet_id", model.ID), zap.Error(err))
		return nil, err
	}

	err = s.tweetService.UpdateLargeImage(model.ID, response.CaptureURL)
	if err != nil {
		s.log.Error("browser CaptureSaveUpdateDatabase, service.UpdateLargeImage", zap.String("tweet_id", model.ID), zap.Error(err))
		return nil, err
	}

	return response, nil
}

func (s serviceImpl) SaveCapture(originalImage []byte, model *service.CaptureRequestModel) (*service.CaptureResponseModel, error) {
	imageKey := fmt.Sprintf("capture/large/%s.jpg", model.ID)

	err := s.bucket.WriteAll(context.Background(), imageKey, originalImage, &blob.WriterOptions{
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
		s.log.Error("browser:saveCapture", zap.String("tweet_id", model.ID), zap.Error(err))
		return nil, err
	}

	return &service.CaptureResponseModel{
		ID:         model.ID,
		CaptureURL: imageKey,
	}, nil
}
func (s serviceImpl) Close() {
	if s.browser != nil {
		s.log.Info("closing browser context")
		s.browser.cancelFunc()
	}
}

func (s *serviceImpl) CaptureURL(model *service.CaptureRequestModel) ([]byte, error) {
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

		allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
		ctx, cancel := chromedp.NewContext(allocCtx)
		s.browser = &browserCtx{
			browserContext: ctx,
			cancelFunc:     cancel,
		}
	}

	var buf []byte
	err := chromedp.Run(s.browser.browserContext, fullScreenshot(model.Url, 90, &buf))
	if err != nil {
		s.log.Error("could not capture URL", zap.String("tweet_id", model.ID), zap.String("url", model.Url), zap.Error(err))
		return nil, err
	}
	return buf, nil
}

func fullScreenshot(url string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(time.Millisecond * 2000),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
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
