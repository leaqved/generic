package set

import "sync"

type Set[T comparable] struct {
	mu   *sync.Mutex
	data map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		mu:   &sync.Mutex{},
		data: make(map[T]struct{}),
	}
}

func (s *Set[T]) Size() int {
	var size int
	s.mu.Lock()
	size = len(s.data)
	s.mu.Unlock()
	return size
}

func (s *Set[T]) Add(elem T) {
	s.mu.Lock()
	s.data[elem] = struct{}{}
	s.mu.Unlock()
}

func (s *Set[T]) Remove(elem T) {
	s.mu.Lock()
	delete(s.data, elem)
	s.mu.Unlock()
}
