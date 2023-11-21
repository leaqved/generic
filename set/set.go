package set

import "sync"

type Set[T comparable] struct {
	Mu   *sync.Mutex
	Data map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		Mu:   &sync.Mutex{},
		Data: make(map[T]struct{}),
	}
}

func (s *Set[T]) Size() int {
	var size int
	s.Mu.Lock()
	size = len(s.Data)
	s.Mu.Unlock()
	return size
}

func (s *Set[T]) Add(elem T) {
	s.Mu.Lock()
	s.Data[elem] = struct{}{}
	s.Mu.Unlock()
}
