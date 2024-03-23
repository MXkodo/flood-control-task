package floodcontrol

import (
	"context"
	"sync"
	"time"
)

type FloodControl interface {
	Check(ctx context.Context, userID int64) (bool, error)
}

type floodControl struct {
	mu   sync.Mutex
	data map[int64][]callInfo
	N    int // Number of seconds to check
	K    int // Maximum number of calls
}

type callInfo struct {
	timestamp time.Time
}

func NewFloodControl(N, K int) FloodControl {
	return &floodControl{
		data: make(map[int64][]callInfo),
		N:    N,
		K:    K,
	}
}

func (fc *floodControl) Check(ctx context.Context, userID int64) (bool, error) {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	now := time.Now()
	for id, calls := range fc.data {
		var newCalls []callInfo
		for _, call := range calls {
			if now.Sub(call.timestamp) < time.Duration(fc.N)*time.Second {
				newCalls = append(newCalls, call)
			}
		}
		fc.data[id] = newCalls
	}

	if len(fc.data[userID]) >= fc.K {
		return false, nil
	}

	fc.data[userID] = append(fc.data[userID], callInfo{timestamp: now})

	return true, nil
}
