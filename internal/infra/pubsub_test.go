package infra

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gocloud.dev/pubsub"
	"testing"
)

func TestNewTopic(t *testing.T) {
	topic, err := NewTopic("mem://test")

	require.NoError(t, err)
	if assert.NotNil(t, topic) {
		err = topic.Send(context.Background(), &pubsub.Message{
			Body: []byte("hello"),
		})
		assert.NoError(t, err)
	}
}
