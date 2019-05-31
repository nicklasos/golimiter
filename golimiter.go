package golimiter

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type Limit struct {
	visitors map[string]*visitor
	mtx      sync.Mutex
	rate     rate.Limit
	bucket   int
}

func NewLimiter(r rate.Limit, bucket int) *Limit {
	limit := &Limit{
		visitors: make(map[string]*visitor),
		rate:     r,
		bucket:   bucket,
	}

	go limit.cleanup()

	return limit
}

func (l *Limit) cleanup() {
	for {
		time.Sleep(time.Minute * 2)
		l.mtx.Lock()
		for id, v := range l.visitors {
			if time.Now().Sub(v.lastSeen) > 5*time.Minute {
				delete(l.visitors, id)
			}
		}
		l.mtx.Unlock()
	}
}

func (l *Limit) addVisitor(id string) *rate.Limiter {
	limiter := rate.NewLimiter(l.rate, l.bucket)
	l.mtx.Lock()

	l.visitors[id] = &visitor{limiter, time.Now()}
	l.mtx.Unlock()

	return limiter
}

func (l *Limit) getVisitor(id string) *rate.Limiter {
	l.mtx.Lock()
	v, exists := l.visitors[id]
	if !exists {
		l.mtx.Unlock()
		return l.addVisitor(id)
	}

	v.lastSeen = time.Now()
	l.mtx.Unlock()

	return v.limiter
}

func (l *Limit) Allow(id string) bool {
	return l.getVisitor(id).Allow()
}
