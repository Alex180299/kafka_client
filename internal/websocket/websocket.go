package websocket

import (
	"context"
	"kafka_client/internal"
)

type WebSocket struct {

}

func (w WebSocket) PublishConsumed(ctx context.Context, message internal.Message) error {
	return nil
}
