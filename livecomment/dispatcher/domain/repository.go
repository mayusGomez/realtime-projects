package domain

type GatewayStorage interface {
	Store(queue, video string)
	Remove(queue, video string)
}

type GatewayConfig interface {
	GetQueues(video string) map[string]struct{}
}
