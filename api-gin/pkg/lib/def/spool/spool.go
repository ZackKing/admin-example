package spool

import (
	"admin-api/pkg/lib/def/slist"
	"sync"
	"time"
)

type SPool struct {
	list        *slist.SList
	mu          sync.RWMutex
	closed      bool
	TTL         time.Duration
	runFunc     RunFunc
	timeoutFunc TimeoutFunc
	ticker      *time.Ticker
}

type RunFunc func() (any, error)
type TimeoutFunc func(any)

type sPoolItem struct {
	val      any
	expireAt int64
}

func New(ttl time.Duration, runFunc RunFunc, timeoutFunc TimeoutFunc) *SPool {
	s := &SPool{
		list:        slist.New(),
		closed:      false,
		TTL:         ttl,
		runFunc:     runFunc,
		timeoutFunc: timeoutFunc,
		ticker:      time.NewTicker(time.Second),
	}
	defer s.ticker.Stop()
	for range s.ticker.C {
		if s.closed {
			break
		}
		s.checkExpired()
	}
	// time.Timer.AddSingleton(context.Background(), time.Second, s.checkExpired)
	return s
}

// Put adds an item to the pool.
func (s *SPool) Put(v any) error {
	if s.closed {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	item := &sPoolItem{
		val: v,
	}
	s.list.Push()
	if s.TTL > 0 {
		item.expireAt = time.Now().Add(s.TTL).Unix()
	} else {
		item.expireAt = 0
	}
	s.list.Push(item)
	return nil
}

// MustPut adds an item to the pool and panics if an error occurs.
func (s *SPool) MustPut(v any) {
	if err := s.Put(v); err != nil {
		panic(err)
	}
}

// Size returns the number of items in the pool.
func (s *SPool) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.list.Len()
}

// Close closes the pool and clears all items.
func (s *SPool) Close() {
	s.closed = true
	s.list.Clear()
}

func (s *SPool) checkExpired() {
	if s.list.Len() == 0 || s.TTL == 0 || s.closed {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().Unix()
	s.list.Each(func(i int, v any) {
		item := v.(*sPoolItem)
		if item.expireAt < now {
			s.timeoutFunc(item.val)
			s.list.Del(i)
		}
	})
}
