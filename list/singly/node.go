package singly

import "sync"

type Node[T any] struct {
	mu    *sync.Mutex
	Value T
	Next  *Node[T]
}

func New[T any]() *Node[T] {
	return &Node[T]{
		mu: &sync.Mutex{},
	}
}

func (n *Node[T]) SetValue(value T) {
	n.mu.Lock()
	n.Value = value
	n.mu.Unlock()
}

func (n *Node[T]) SetNext(next *Node[T]) {
	n.mu.Lock()
	n.Next = next
	n.mu.Unlock()
}

func (n *Node[T]) LoadValue() T {
	var value T
	n.mu.Lock()
	value = n.Value
	n.mu.Unlock()
	return value
}

func (n *Node[T]) LoadNext() *Node[T] {
	var next *Node[T]
	n.mu.Lock()
	next = n.Next
	n.mu.Unlock()
	return next
}

func (n *Node[T]) SwapValue(value T) T {
	n.mu.Lock()
	value, n.Value = n.Value, value
	n.mu.Unlock()
	return value
}

func (n *Node[T]) SwapNext(next *Node[T]) *Node[T] {
	n.mu.Lock()
	next, n.Next = n.Next, next
	n.mu.Unlock()
	return next
}
