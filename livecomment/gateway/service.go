package gateway

import (
	"github.com/gin-gonic/gin"
	"livecomments/gateway/application"
	"livecomments/gateway/interfaces/web"
)

type ServiceContainer struct {
	SubscriptionHandler func(c *gin.Context)
}

func NewService() *ServiceContainer {
	subsSvc := application.NewSubscriptionService()
	subsHandler := web.NewSubscriptionHandler(subsSvc.Subscribe, subsSvc.Unsubscribe)

	return &ServiceContainer{
		SubscriptionHandler: subsHandler.Handle,
	}
}
