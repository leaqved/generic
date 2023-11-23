package singly

import "sync"

type Node[T any] struct {
	Mu   *sync.Mutex
	Val  T
	Next *Node[T]
}

func New[T any]() *Node[T] {
	return &Node[T]{
		Mu: &sync.Mutex{},
	}
}
