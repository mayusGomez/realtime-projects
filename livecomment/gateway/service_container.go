package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"livecomments/gateway/application"
	"livecomments/gateway/interfaces/rabbit"
	"livecomments/gateway/interfaces/web"
	"livecomments/pkg/adapters"
)

type ServiceContainer struct {
	SubscriptionHandler func(c *gin.Context)
	consumerHandler     func(rabbitMsg *amqp091.Delivery) error
}

func NewService() *ServiceContainer {
	subsSvc := application.NewSubscriptionService()
	subsHandler := web.NewSubscriptionHandler(subsSvc.Subscribe, subsSvc.Unsubscribe)
	consumerHandler := rabbit.NewCommentsHandler(subsSvc)

	return &ServiceContainer{
		SubscriptionHandler: subsHandler.Handle,
		consumerHandler:     consumerHandler.Handle,
	}
}

func (sc *ServiceContainer) Run(port, rabbitConnStr, queue string) error {
	appAdapter := adapters.NewAppAdapters()
	appAdapter.AddAdapters(
		NewWebAdapter(port, sc.SubscriptionHandler),
		NewRabbitMQAdapter(rabbitConnStr, queue, sc.consumerHandler),
	)

	appAdapter.Run()

	return nil
}
