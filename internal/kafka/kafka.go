package kafka

import (
	"context"
	"kafka_client/internal"
)

type Kafka struct {

}

func (k Kafka) Publish(ctx context.Context, message internal.Message) error {
	return nil
}

func (k Kafka) Listen(ctx context.Context, topic string) error {
	return nil
}

/*
func NewKafka(newEvent *external_events.NewEvent, address string) Kafka {
	return Kafka{
		newEvent: newEvent,
		address:  address,
	}
}

func (i *Kafka) Produce(event application.Event) error {
	log.Println(fmt.Sprintf("producing new event %+v", event))
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", i.address, string(event.Topic), partition)
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
		kafka.Message{Key: []byte(event.Key), Value: []byte(event.Value)},
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

func (i *Kafka) Listen(address string, topic application.Topic) error {
	reader := getKafkaReader(address, string(topic))
	err := reader.SetOffset(1)
	if err != nil {
		return err
	}
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		err = i.newEvent.Consume(application.Event{
			Key:   string(m.Key),
			Value: string(m.Value),
			Topic: application.Topic(m.Topic),
		})
		if err != nil {
			log.Fatal(fmt.Sprintf("Error to consume new event: (%+v), %s", m, err.Error()))
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
*/