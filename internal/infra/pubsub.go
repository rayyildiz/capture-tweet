package infra

import (
	"context"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/gcppubsub"
	_ "gocloud.dev/pubsub/mempubsub"
	"os"
)

func NewTopic(topic string) (*pubsub.Topic, error) {
	return pubsub.OpenTopic(context.Background(), topic)
}

func NewCaptureTopic() (*pubsub.Topic, error) {
	return NewTopic(os.Getenv("TOPIC_CAPTURE"))
}
