package domain

type SubscriptionCmd interface {
	Subscribe(queue, video string) error
	Unsubscribe(queue, video string) error
}

type PostCommentCmd interface {
	PostComment(connectionId, video, comment string) error
}
