package application

import "livecomments/dispatcher/domain"

type SubscribeGateway struct {
	gatewayStorage domain.GatewayStorage
}

func NewSubscribeGateway(gatewayStorage domain.GatewayStorage) *SubscribeGateway {
	return &SubscribeGateway{
		gatewayStorage: gatewayStorage,
	}
}

func (g *SubscribeGateway) Subscribe(queue, video string) error {
	g.gatewayStorage.Store(queue, video)

	return nil
}

func (g *SubscribeGateway) Unsubscribe(queue, video string) error {
	g.gatewayStorage.Remove(queue, video)

	return nil
}
