package linear

import (
	s "generic/list/singly"
	"sync"
)

type List[T any] struct {
	mu     *sync.Mutex
	head   *s.Node[T]
	length int
}

func New[T any]() *List[T] {
	return &List[T]{
		mu: &sync.Mutex{},
	}
}

func (l *List[T]) Add(value T) {
	l.mu.Lock()
	head := s.New[T]()
	head.Value = value
	head.Next = l.head
	l.head = head
	l.length++
	l.mu.Unlock()
}

func (l *List[T]) Remove() {
	l.mu.Lock()
	head := l.head.Next
	l.head = head
	l.length--
	l.mu.Unlock()
}

func (l *List[T]) Len() int {
	var len int
	l.mu.Lock()
	len = l.length
	l.mu.Unlock()
	return len
}
