package dispatcher

import (
	"github.com/gin-gonic/gin"
	"livecomments/dispatcher/application"
	"livecomments/dispatcher/infrastructure/gatewayconfig"
	rabbitmqpublisher "livecomments/dispatcher/infrastructure/rabbitmq"
	"livecomments/dispatcher/interfaces/web"
	"livecomments/pkg/adapters"
	"livecomments/pkg/rabbitmq"
)

type ServiceContainer struct {
	subscriptionHandler func(c *gin.Context)
	postCommentHandler  func(c *gin.Context)
}

func NewServiceContainer(rabbitConnStr string, queues []string) (*ServiceContainer, error) {
	gatewayStorage := gatewayconfig.NewStorage()
	subsSvc := application.NewSubscribeGateway(gatewayStorage)
	subsHandler := web.NewSubscribeGateway(subsSvc)

	rabbitClient, err := rabbitmq.NewRabbitClient(rabbitConnStr, queues)
	if err != nil {
		return nil, err
	}

	publisher := rabbitmqpublisher.NewPublisher(rabbitClient)
	postCommentCmd := application.NewComment(gatewayStorage, publisher)
	commentHandler := web.NewCommentHandler(postCommentCmd)

	return &ServiceContainer{
		subscriptionHandler: subsHandler.Handle,
		postCommentHandler:  commentHandler.Handle,
	}, nil
}

func (sc *ServiceContainer) Run(port string) error {
	appAdapter := adapters.NewAppAdapters()
	appAdapter.AddAdapters(
		NewWebAdapter(port, sc.subscriptionHandler, sc.postCommentHandler),
	)

	appAdapter.Run()

	return nil
}
