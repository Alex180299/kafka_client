package interfaces

import "kafka_client/application"

type ReportEventRepository interface {
	ReportEvent(event application.Event) error
}

type PublishEventRepository interface {
	PublishEvent(event application.Event) error
}

type ListenTopicRepository interface {
	ListenTopic(topic application.Topic) error
}
