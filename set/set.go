// Package set provides generic Set data structure and its basic methods.
package set

import "sync"

// Set is a data structure that can store unique values, without any particular order.
//
// Set consists of elements of type T.
//
// Set is thread-safe.
type Set[T comparable] struct {
	mu   *sync.Mutex
	data map[T]struct{}
}

// New creates a Set of type T.
func New[T comparable]() *Set[T] {
	return &Set[T]{
		mu:   &sync.Mutex{},
		data: make(map[T]struct{}),
	}
}

// Size returns the number of elements in the Set.
func (s *Set[T]) Size() int {
	var size int
	s.mu.Lock()
	size = len(s.data)
	s.mu.Unlock()
	return size
}

// Add adds the element to the Set.
func (s *Set[T]) Add(elem T) {
	s.mu.Lock()
	s.data[elem] = struct{}{}
	s.mu.Unlock()
}

// Remove removes the element from the Set.
func (s *Set[T]) Remove(elem T) {
	s.mu.Lock()
	delete(s.data, elem)
	s.mu.Unlock()
}

func (s *Set[T]) Contains(elem T) bool {
	var res bool
	s.mu.Lock()
	_, res = s.data[elem]
	s.mu.Unlock()
	return res
}

func (s *Set[T]) Clear() {
	s.mu.Lock()
	for elem := range s.data {
		delete(s.data, elem)
	}
	s.mu.Unlock()
}
