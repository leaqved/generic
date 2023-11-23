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
