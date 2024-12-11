package rabbitmq

import (
	"encoding/json"
	"livecomments/dispatcher/domain"
	"livecomments/pkg/rabbitmq"
)

type Publisher struct {
	rabbitClient *rabbitmq.RabbitMQClient
}

func NewPublisher(rabbitClient *rabbitmq.RabbitMQClient) *Publisher {
	return &Publisher{rabbitClient: rabbitClient}
}

func (p *Publisher) PostMessage(queues map[string]struct{}, comment *domain.CommentMessage) error {
	commentByte, err := json.Marshal(comment)
	if err != nil {
		return err
	}

	for queue := range queues {
		err = p.rabbitClient.Publish(queue, commentByte)
		if err != nil {
			return err
		}
	}

	return nil
}
