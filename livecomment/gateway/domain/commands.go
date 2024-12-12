package domain

type CommentPublisher interface {
	PublishComment(video, message string) error
}
