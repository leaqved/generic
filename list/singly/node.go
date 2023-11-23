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

func (n *Node[T]) SetVal(val T) {
	n.Mu.Lock()
	n.Val = val
	n.Mu.Unlock()
}
