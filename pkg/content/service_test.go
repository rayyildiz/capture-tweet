package content

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestService_StoreContactRequest(t *testing.T) {
	os.Setenv("CAPTCHA_SECRET", "6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe")
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().ContactUs(ctx, "test@eample.com", "user", "hello").Return(nil)

	svc := NewService(repo)
	require.NotNil(t, svc)

	err := svc.StoreContactRequest(ctx, "test@eample.com", "user", "hello", "6LeIxAcTAAAAAJcZVRqyHh71UMIEGNQ_MXjiZKhI")
	require.NoError(t, err)
}
