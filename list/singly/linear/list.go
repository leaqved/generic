package linear

import (
	s "generic/list/singly"
	"sync"
)

type List[T any] struct {
	Mu     *sync.Mutex
	Head   *s.Node[T]
	Length int
}

func New[T any]() *List[T] {
	return &List[T]{
		Mu: &sync.Mutex{},
	}
}

func (l *List[T]) Add(val T) {
	l.Mu.Lock()
	head := s.New[T]()
	head.Val = val
	head.Next = l.Head
	l.Head = head
	l.Length++
	l.Mu.Unlock()
}

func (l *List[T]) Remove() {
	l.Mu.Lock()
	head := l.Head.Next
	l.Head = head
	l.Length--
	l.Mu.Unlock()
}
