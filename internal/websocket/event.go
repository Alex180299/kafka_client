package websocket

import (
	"context"
	"encoding/json"
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
	"kafka_client/internal"
)

type EventWebSocket struct {
	connections []*socketio.Conn
}

func NewEventWebSocket() EventWebSocket {
	return EventWebSocket{
		connections: make([]*socketio.Conn, 0),
	}
}

func (w *EventWebSocket) Register(echoServer *echo.Echo) {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(conn socketio.Conn) error {
		w.connections = append(w.connections, &conn)
		return nil
	})
	server.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		w.connections = make([]*socketio.Conn, 0)

		for _, c := range w.connections {
			if (*c).ID() != conn.ID() {
				w.connections = append(w.connections, &conn)
			}
		}
	})

	go server.Serve()

	echoServer.Any("/socket.io/", func(c echo.Context) error {
		server.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}

func (w *EventWebSocket) PublishConsumed(ctx context.Context, message internal.Message) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}

	for _, c := range w.connections {
		(*c).Emit("/", msg)
	}

	return nil
}
