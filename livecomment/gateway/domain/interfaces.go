package domain

type DispatcherSubscriber interface {
	Subscribe(video, queue string) error
	Unsubscribe(video, queue string) error
}
