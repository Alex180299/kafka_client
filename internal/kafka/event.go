package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"kafka_client/internal"
	"log"
	"strings"
	"time"
)

type ConsumerService interface {
	Consume(ctx context.Context, message internal.Message) error
}

type EventKafka struct {
	svc ConsumerService
}

func (k EventKafka) Publish(ctx context.Context, message internal.Message) error {
	log.Println(fmt.Sprintf("producing new event %+v", message))
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", message.Address, string(message.Topic), partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
		return err
	}

	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return err
	}
	_, err = conn.WriteMessages(
		kafka.Message{Key: []byte(message.Key), Value: []byte(message.Value)},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return err
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
		return err
	}

	return nil
}

func (k EventKafka) Listen(ctx context.Context, topic string, address string) error {
	reader := getKafkaReader(address, topic)

	err := reader.SetOffset(1)
	if err != nil {
		return err
	}

	go k.listenInBackground(reader)

	return nil
}

func (k EventKafka) listenInBackground(reader *kafka.Reader) {
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		err = k.svc.Consume(context.Background(), internal.Message{
			Key:   string(m.Key),
			Value: string(m.Value),
			Topic: internal.Topic(m.Topic),
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getKafkaReader(address, topic string) *kafka.Reader {
	brokers := strings.Split(address, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MinBytes: 10e3,  // 10KB
		MaxBytes: 100e6, // 10MB
	})
}
