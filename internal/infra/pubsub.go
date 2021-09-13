package infra

import (
	"context"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/gcppubsub"
	_ "gocloud.dev/pubsub/mempubsub"
	"log"
	// _ "gocloud.dev/pubsub/natspubsub"
)

func NewTopic(topicName string) *pubsub.Topic {
	topic, err := pubsub.OpenTopic(context.Background(), topicName)
	if err != nil {
		log.Fatalf("error while opening pubsub topic, %v", err)
	}
	return topic
}
