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
