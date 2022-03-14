package infra

import (
	"context"
	"github.com/matryer/is"
	"gocloud.dev/pubsub"
	"testing"
)

func TestNewTopic(t *testing.T) {
	is := is.New(t)

	topic := NewTopic("mem://test")

	is.True(nil != topic)
	err := topic.Send(context.Background(), &pubsub.Message{
		Body: []byte("hello"),
	})
	is.NoErr(err)
}
