// Package linear provides generic Singly Linked List (SLL) data structure and its basic methods.
package linear

import (
	s "generic/list/singly"
	"sync"
)

// Singly Linked List (SLL) is a linear data structure that can store values.
//
// SLL consists of Singly Linked Nodes (SLN) of type T.
//
// SLL is thread-safe.
type List[T any] struct {
	mu     *sync.Mutex
	head   *s.Node[T]
	length int
}

// New creates a List of type T.
func New[T any]() *List[T] {
	return &List[T]{
		mu: &sync.Mutex{},
	}
}

// Add inserts the value at the beginning of the List.
func (l *List[T]) Add(value T) {
	l.mu.Lock()
	head := s.New[T]()
	head.Value = value
	head.Next = l.head
	l.head = head
	l.length++
	l.mu.Unlock()
}

// Removes the first Node from the List.
func (l *List[T]) Remove() {
	l.mu.Lock()
	head := l.head.Next
	l.head = head
	l.length--
	l.mu.Unlock()
}

// Len returns the number of elements in the List.
func (l *List[T]) Len() int {
	var len int
	l.mu.Lock()
	len = l.length
	l.mu.Unlock()
	return len
}
