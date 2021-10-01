package domain

import (
	"kafka_client/application"
	"kafka_client/domain/interfaces"
)

type Event struct {
	reportEventRepository  interfaces.ReportEventRepository
	publishEventRepository interfaces.PublishEventRepository
	listenTopicRepository  interfaces.ListenTopicRepository
}

func NewEventDomain(reportEventRepository interfaces.ReportEventRepository, publishEventRepository interfaces.PublishEventRepository,
	listenTopicRepository interfaces.ListenTopicRepository) *Event {
	return &Event{
		reportEventRepository:  reportEventRepository,
		publishEventRepository: publishEventRepository,
		listenTopicRepository:  listenTopicRepository,
	}
}

func (e *Event) publish(event application.Event) error {
	err := e.publishEventRepository.PublishEvent(event)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) consume(event application.Event) error {
	err := e.reportEventRepository.ReportEvent(event)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) listenTopic(topic application.Topic) error {
	err := e.listenTopicRepository.ListenTopic(topic)
	if err != nil {
		return err
	}

	return nil
}
