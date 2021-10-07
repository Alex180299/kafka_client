package service

import (
	"context"
	"kafka_client/internal"
)

type Socket interface {
	PublishConsumed(ctx context.Context, message internal.Message) error
}

type Consumer struct {
	socket Socket
}

func NewConsumer(socket Socket) Consumer {
	return Consumer{
		socket: socket,
	}
}

func (c *Consumer) Consume(ctx context.Context, message internal.Message) error {
	if err := message.Validate(); err != nil {
		return err
	}

	err := c.socket.PublishConsumed(ctx, message)
	if err != nil {
		return err
	}

	return nil
}
