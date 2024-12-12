package application

import (
	"errors"
	"livecomments/gateway/domain"
	"log"
	"sync"
)

type SubscriptionService struct {
	mu                sync.RWMutex
	videoSubscription map[string]map[string]chan string
	dispatcher        domain.DispatcherSubscriber
	queue             string
}

func NewSubscriptionService(dispatcher domain.DispatcherSubscriber, queue string) *SubscriptionService {
	return &SubscriptionService{
		videoSubscription: make(map[string]map[string]chan string),
		dispatcher:        dispatcher,
		queue:             queue,
	}
}

func (s *SubscriptionService) Subscribe(video, connectionId string) (chan string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	videoSubscription, ok := s.videoSubscription[video]
	if !ok {
		err := s.dispatcher.Subscribe(video, s.queue)
		if err != nil {
			return nil, err
		}
		videoSubscription = make(map[string]chan string)
	}

	if _, ok := videoSubscription[connectionId]; ok {
		return nil, errors.New("already subscribed")
	}

	ch := make(chan string, 10)
	videoSubscription[connectionId] = ch
	log.Println("subscribe to video and connection id:", video, connectionId)

	s.videoSubscription[video] = videoSubscription

	return ch, nil
}

func (s *SubscriptionService) Unsubscribe(video, connectionId string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	videoSubscription, ok := s.videoSubscription[video]
	if !ok {
		return
	}

	// TODO: include unsubscribe video from dispatcher

	connection, ok := videoSubscription[connectionId]
	if !ok {
		return
	}

	close(connection)
	delete(videoSubscription, connectionId)
}

func (s *SubscriptionService) PublishComment(video, message string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, client := range s.videoSubscription[video] {
		select {
		case client <- message:
		default:
			log.Println("cannot send message")
		}
	}

	return nil
}
