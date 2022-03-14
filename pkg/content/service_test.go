package content

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/matryer/is"
	"os"
	"testing"
)

func TestService_StoreContactRequest(t *testing.T) {
	is := is.New(t)

	os.Setenv("CAPTCHA_SECRET", "6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().ContactUs(gomock.Any(), "test@eample.com", "user", "hello").Return(nil)

	svc := NewService(repo)
	is.True(svc != nil)

	err := svc.StoreContactRequest(context.Background(), "test@eample.com", "user", "hello", "6LeIxAcTAAAAAJcZVRqyHh71UMIEGNQ_MXjiZKhI")
	is.NoErr(err)
}
