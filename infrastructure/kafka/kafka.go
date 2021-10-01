package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
	"time"
)

func Consume(topic string) {
	log.Println("consumiendo")
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafkacluster.dev.rappi.com:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// conn.SetReadDeadline(time.Now().Add(10*time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 100e3) // 10KB max per message
	for {
		_, err := conn.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

func Produce(topic string, value []byte){
	log.Println("Produciendo")
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafkacluster.dev.rappi.com:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	err = conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	_, err = conn.WriteMessages(
		kafka.Message{Key: []byte("key"), Value: value},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	log.Println("Terminando de producir")
}

func Listen(topic string) {
	reader := getKafkaReader(topic)
	err := reader.SetOffset(1)
	if err != nil {
		return
	}
	defer reader.Close()


	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		// fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		fmt.Printf("%s\n", string(m.Value))
	}
}

func getKafkaReader(topic string) *kafka.Reader {
	brokers := strings.Split("kafkacluster.dev.rappi.com:9092", ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 100e6, // 10MB
	})
}
