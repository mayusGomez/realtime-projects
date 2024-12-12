package gatewayconfig

import (
	"reflect"
	"sync"
	"testing"
)

func TestStorage_GetQueues(t *testing.T) {
	data := map[string]map[string]struct{}{
		"video-001": {
			"queue-001": struct{}{},
			"queue-002": struct{}{},
		},
		"video-002": {},
	}
	type fields struct {
		mu   sync.RWMutex
		data map[string]map[string]struct{}
	}
	type args struct {
		video string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]struct{}
	}{
		{
			name: "Get all queues with return",
			fields: fields{
				data: data,
			},
			args: args{
				video: "video-001",
			},
			want: map[string]struct{}{
				"queue-001": struct{}{},
				"queue-002": struct{}{},
			},
		},
		{
			name: "Get all queues without return",
			fields: fields{
				data: data,
			},
			args: args{
				video: "video-002",
			},
			want: map[string]struct{}{},
		},
		{
			name: "Get video not exist",
			fields: fields{
				data: data,
			},
			args: args{
				video: "video-003",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				mu:   tt.fields.mu,
				data: tt.fields.data,
			}
			if got := s.GetQueues(tt.args.video); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQueues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStorage(t *testing.T) {
	storage := NewStorage()

	if storage == nil {
		t.Fatal("Expected storage to be initialized, got nil")
	}

	if len(storage.data) != 0 {
		t.Fatalf("Expected storage.data to be empty, got length: %d", len(storage.data))
	}
}

func TestStore(t *testing.T) {
	storage := NewStorage()

	storage.Store("queue1", "video1")
	storage.Store("queue2", "video1")

	if len(storage.data) != 1 {
		t.Fatalf("Expected 1 video in storage, got %d", len(storage.data))
	}

	videoSet, ok := storage.data["video1"]
	if !ok {
		t.Fatalf("Expected video1 to exist in storage")
	}

	if len(videoSet) != 2 {
		t.Fatalf("Expected 2 queues for video1, got %d", len(videoSet))
	}

	if _, exists := videoSet["queue1"]; !exists {
		t.Fatalf("Expected queue1 to exist for video1")
	}

	if _, exists := videoSet["queue2"]; !exists {
		t.Fatalf("Expected queue2 to exist for video1")
	}
}

func TestRemove(t *testing.T) {
	storage := NewStorage()

	storage.Store("queue1", "video1")
	storage.Store("queue2", "video1")
	storage.Remove("queue1", "video1")

	videoSet, ok := storage.data["video1"]
	if !ok {
		t.Fatalf("Expected video1 to exist in storage")
	}

	if len(videoSet) != 1 {
		t.Fatalf("Expected 1 queue for video1 after removal, got %d", len(videoSet))
	}

	if _, exists := videoSet["queue1"]; exists {
		t.Fatalf("Did not expect queue1 to exist for video1 after removal")
	}

	if _, exists := videoSet["queue2"]; !exists {
		t.Fatalf("Expected queue2 to still exist for video1")
	}
}

func TestRemoveNonExistent(t *testing.T) {
	storage := NewStorage()

	storage.Store("queue1", "video1")
	storage.Remove("queue2", "video1") // Removing a non-existent queue
	storage.Remove("queue1", "video2") // Removing from a non-existent video

	videoSet, ok := storage.data["video1"]
	if !ok {
		t.Fatalf("Expected video1 to exist in storage")
	}

	if len(videoSet) != 1 {
		t.Fatalf("Expected 1 queue for video1 after non-existent removals, got %d", len(videoSet))
	}

	if _, exists := videoSet["queue1"]; !exists {
		t.Fatalf("Expected queue1 to still exist for video1")
	}
}
