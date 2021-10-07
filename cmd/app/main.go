package main

import (
	"github.com/labstack/echo/v4"
	"kafka_client/internal/kafka"
	"kafka_client/internal/rest"
	"kafka_client/internal/service"
	"kafka_client/internal/websocket"
)

func main() {
	newServer()
}

func newServer() {
	e := echo.New()

	socket := websocket.NewEventWebSocket()
	consumerSvc := service.NewConsumer(&socket)

	broker := kafka.NewEventKafka(&consumerSvc)
	publisherSvc := service.NewPublisher(broker)

	eventHandler := rest.NewEventHandler(&publisherSvc)
	eventHandler.Register(e)
	socket.Register(e)

	e.Static("/", "ui")
	e.Logger.Fatal(e.Start(":8080"))
}
