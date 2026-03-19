package main

import (
	"fmt"
	"sync"
	"time"
)

type rateLimit struct {
	capacity  int
	refill    int
	lastrefil time.Time
	tokens    int
	mu        sync.Mutex
}

func NewRateLimiter(cap int, refil int) *rateLimit {
	return &rateLimit{
		capacity:  cap,
		refill:    refil,
		tokens:    cap,
		lastrefil: time.Now(),
	}
}

func (r *rateLimit) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Refill()
	if r.tokens > 0 {
		r.tokens--
		return true
	}
	return false
}

func (r *rateLimit) Refill() {
	now := time.Now()
	refil := now.Sub(r.lastrefil).Seconds()
	tokenstoAdd := int(refil * float64(r.refill))

	if tokenstoAdd >= r.capacity {
		tokenstoAdd = r.capacity
	}
	if tokenstoAdd != 0 {
		r.tokens = tokenstoAdd
	}
	r.lastrefil = time.Now()
}

func main() {
	rateLimiter := NewRateLimiter(5, 2)
	fmt.Println(rateLimiter)
	for i := 0; i <= 10; i++ {
		if rateLimiter.Allow() {
			fmt.Println("allow")
		} else {
			fmt.Println("do not allow")
		}
		if i == 7 {
			time.Sleep(2 * time.Second)
		}
	}
}
