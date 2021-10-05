package rest

import (
	"context"
	"github.com/labstack/echo/v4"
	"kafka_client/internal"
	"net/http"
)

type PublisherService interface {
	Publish(ctx context.Context, message internal.Message) (internal.Message, error)
	Listen(ctx context.Context, topic internal.Topic) error
}

type EventHandler struct {
	svc PublisherService
}

func (e EventHandler) Register(echoServer *echo.Echo) {
	echoServer.POST("/listen", e.listen)
	echoServer.POST("/publish", e.publish)
}

type listenRequest struct {
	Topic string `json:"topic"`
}

func (e *EventHandler) listen(c echo.Context) error {
	req := new(listenRequest)
	err := c.Bind(req)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err = e.svc.Listen(c.Request().Context(), internal.Topic(req.Topic))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

type publishRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Topic string `json:"topic"`
}

func (e *EventHandler) publish(c echo.Context) error {
	req := new(publishRequest)
	err := c.Bind(req)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	_, err = e.svc.Publish(c.Request().Context(), internal.Message{
		Key:   req.Key,
		Value: req.Value,
		Topic: internal.Topic(req.Topic),
	})
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
