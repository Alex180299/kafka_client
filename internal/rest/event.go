package rest

import (
	"context"
	"kafka_client/internal"
)

type PublisherService interface {
	Publish(ctx context.Context, message internal.Message) (internal.Message, error)
	Listen(ctx context.Context, topic internal.Topic) error
}

type EventHandler struct {
	svc PublisherService
}

func (e EventHandler) Register() {

}
