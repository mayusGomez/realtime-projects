package gateway

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"livecomments/pkg/rabbitmq"
	"log"
)

type RabbitMQAdapter struct {
	connStr            string
	queueName          string
	newCommentsHandler func(rabbitMsg *amqp091.Delivery) error
	rabbitClient       *rabbitmq.RabbitMQClient
}

func NewRabbitMQAdapter(connStr string, queueName string, newCommentConsumerFn func(rabbitMsg *amqp091.Delivery) error) *RabbitMQAdapter {
	return &RabbitMQAdapter{
		newCommentsHandler: newCommentConsumerFn,
		connStr:            connStr,
		queueName:          queueName,
	}
}

func (r *RabbitMQAdapter) Start(_ context.Context) error {
	rb, err := rabbitmq.NewRabbitClient(r.connStr, []string{r.queueName})
	if err != nil {
		return err
	}

	r.rabbitClient = rb

	msgs, err := rb.SubscribeToQueue(r.queueName)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			err = r.newCommentsHandler(&d)
			if err != nil {
				// Don't handle retries, only log the error
				log.Println(err)
			}
		}
	}()

	return nil
}

func (r *RabbitMQAdapter) Stop(_ context.Context) error {
	r.rabbitClient.Close()

	return nil
}
