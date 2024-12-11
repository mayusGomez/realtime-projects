package gateway

import (
	"github.com/gin-gonic/gin"
	"livecomments/gateway/application"
	"livecomments/gateway/interfaces/web"
	"livecomments/pkg/adapters"
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

func (sc *ServiceContainer) Run(port string) error {
	appAdapter := adapters.NewAppAdapters()
	appAdapter.AddAdapters(
		NewWebAdapter(port, sc.SubscriptionHandler),
	)

	appAdapter.Run()

	return nil
}
