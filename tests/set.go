package tests

import "sync"

type Set struct {
	set map[string]struct{}
	mu  sync.Mutex
}

func (s *Set) Add(x string) {
	s.mu.Lock()
	s.set[x] = struct{}{}
	s.mu.Unlock()
}

func (s *Set) Delete(x string) {
	s.mu.Lock()
	delete(s.set, x)
	s.mu.Unlock()
}
