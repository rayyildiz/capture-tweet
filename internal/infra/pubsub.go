package infra

import (
	"context"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/gcppubsub"
	_ "gocloud.dev/pubsub/mempubsub"
	_ "gocloud.dev/pubsub/natspubsub"
)

func NewTopic(topic string) (*pubsub.Topic, error) {
	return pubsub.OpenTopic(context.Background(), topic)
}
