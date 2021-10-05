package websocket

import (
	"context"
	"kafka_client/internal"
)

type EventWebSocket struct {

}

func (w EventWebSocket) PublishConsumed(ctx context.Context, message internal.Message) error {
	return nil
}
