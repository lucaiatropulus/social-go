package ratelimiter

import (
	"github.com/lucaiatropulus/social/internal/utils"
	"log"
	"sync"
	"time"
)

type FixedWindowRateLimiter struct {
	sync.RWMutex
	clients map[string]int
	limit   int
	window  time.Duration
}

func NewFixedWindowRateLimiter(limit int, window string) *FixedWindowRateLimiter {
	windowDuration, ok := utils.ParseStringToDuration(window)

	if !ok {
		log.Fatal("Unable to parse window duration")
	}

	return &FixedWindowRateLimiter{
		clients: make(map[string]int),
		limit:   limit,
		window:  windowDuration,
	}
}

func (l *FixedWindowRateLimiter) Allow(ip string) (bool, time.Duration) {
	l.RLock()
	count, exists := l.clients[ip]
	l.RUnlock()

	if !exists || count < l.limit {
		l.Lock()
		if !exists {
			go l.resetCount(ip)
		}

		l.clients[ip]++
		l.Unlock()
		return true, 0
	}

	return false, l.window
}

func (l *FixedWindowRateLimiter) resetCount(ip string) {
	time.Sleep(l.window)
	l.Lock()
	delete(l.clients, ip)
	l.Unlock()
}
