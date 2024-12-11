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
