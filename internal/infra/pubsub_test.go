package infra

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gocloud.dev/pubsub"
	"testing"
)

func TestNewTopic(t *testing.T) {
	topic := NewTopic("mem://test")

	if assert.NotNil(t, topic) {
		err := topic.Send(context.Background(), &pubsub.Message{
			Body: []byte("hello"),
		})
		assert.NoError(t, err)
	}
}
