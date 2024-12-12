package rabbit

import (
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"livecomments/gateway/domain"
)

type CommentMessage struct {
	ConnectionId string
	Video        string
	Message      string
}

type Comments struct {
	publishMsg domain.CommentPublisher
}

func NewCommentsHandler(publishMsg domain.CommentPublisher) *Comments {
	return &Comments{
		publishMsg: publishMsg,
	}
}

func (c *Comments) Handle(rabbitMsg *amqp091.Delivery) error {
	msg := CommentMessage{}
	err := json.Unmarshal(rabbitMsg.Body, &msg)
	if err != nil {
		return err
	}

	err = c.publishMsg.PublishComment(msg.Video, msg.Message)
	if err != nil {
		return err
	}

	return nil
}
