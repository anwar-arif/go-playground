package rate_limiter

import (
	"sync"
	"time"
)

type Request struct {
	UserId    string
	Timestamp time.Time
}

type RateLimiter struct {
	mu      sync.RWMutex
	Window  time.Duration
	Count   int64
	Records map[string][]Request
}

func NewRateLimiter(count int64, window time.Duration) *RateLimiter {
	return &RateLimiter{
		Window:  window,
		Count:   count,
		Records: make(map[string][]Request),
	}
}

func (rl *RateLimiter) AllowRequest(user string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	curTime := time.Now()
	windowStart := curTime.Add(-rl.Window)

	valid := rl.Records[user][:0]
	for _, req := range rl.Records[user] {
		if req.Timestamp.After(windowStart) {
			valid = append(valid, req)
		}
	}

	rl.Records[user] = valid
	if int64(len(rl.Records[user])) >= rl.Count {
		return false
	}

	rl.Records[user] = append(rl.Records[user], Request{user, curTime})
	return true
}
