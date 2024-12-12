package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"livecomments/gateway/application"
	"livecomments/gateway/infrastructure/dispatcher"
	"livecomments/gateway/interfaces/rabbit"
	"livecomments/gateway/interfaces/web"
	"livecomments/pkg/adapters"
)

type ServiceContainer struct {
	SubscriptionHandler func(c *gin.Context)
	consumerHandler     func(rabbitMsg *amqp091.Delivery) error
}

func NewService(dispatcherURL, queue string) *ServiceContainer {
	dispatcherClient := dispatcher.NewDispatcher(dispatcherURL)
	subsSvc := application.NewSubscriptionService(dispatcherClient, queue)
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
