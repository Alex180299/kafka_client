package service

import (
	"context"
	"kafka_client/internal"
)

type MessageBroker interface {
	Publish(ctx context.Context, message internal.Message) error
	Listen(ctx context.Context, topic string) error
}

type Publisher struct {
	broker MessageBroker
}

func (p *Publisher) Publish(ctx context.Context, message internal.Message) (internal.Message, error) {
	if err := message.Validate(); err != nil {
		return internal.Message{}, err
	}

	err := p.broker.Publish(ctx, message)
	if err != nil {
		return internal.Message{}, err
	}

	return message, nil
}

func (p *Publisher) Listen(ctx context.Context, topic internal.Topic) error {
	if err := topic.Validate(); err != nil {
		return err
	}

	err := p.broker.Listen(ctx, string(topic))
	if err != nil {
		return err
	}

	return nil
}
