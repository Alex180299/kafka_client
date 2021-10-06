package internal

import (
	"errors"
	"time"
)

type Topic string

func (t Topic) Validate() error {
	if t == "" {
		return errors.New("empty topic")
	}

	return nil
}

type Message struct {
	Sent    bool
	Id      uint32
	Key     string
	Value   string
	Address string
	Topic   Topic
	SentAt  time.Time
}

func (m *Message) Validate() error {
	if m.Key == "" {
		return errors.New("empty key")
	}
	if m.Value == "" {
		return errors.New("empty value")
	}
	if m.Address == "" {
		return errors.New("empty address")
	}
	if err := m.Topic.Validate(); err != nil {
		return err
	}

	return nil
}
