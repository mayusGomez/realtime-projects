package gatewayconfig

import "sync"

type Storage struct {
	mu   sync.RWMutex
	data map[string]map[string]struct{}
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]map[string]struct{}),
	}
}

func (s *Storage) Store(queue, video string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	videoSet, ok := s.data[video]
	if !ok {
		s.data[video] = make(map[string]struct{})
	}

	videoSet[queue] = struct{}{}
}

func (s *Storage) Remove(queue, video string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	videoSet, ok := s.data[video]
	if !ok {
		return
	}

	delete(videoSet, queue)
}

func (s *Storage) GetQueues(video string) map[string]struct{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data[video]
}
