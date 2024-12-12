package domain

type CommentMessage struct {
	ConnectionId string
	Video        string
	Message      string
}

type AsyncCommunication interface {
	PostMessage(queues map[string]struct{}, comment *CommentMessage) error
}
