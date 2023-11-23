package singly

import "sync"

type Node[T any] struct {
	mu   *sync.Mutex
	Val  T
	Next *Node[T]
}

func New[T any]() *Node[T] {
	return &Node[T]{
		mu: &sync.Mutex{},
	}
}

func (n *Node[T]) SetVal(val T) {
	n.mu.Lock()
	n.Val = val
	n.mu.Unlock()
}

func (n *Node[T]) SetNext(next *Node[T]) {
	n.mu.Lock()
	n.Next = next
	n.mu.Unlock()
}

func (n *Node[T]) LoadVal() T {
	var val T
	n.mu.Lock()
	val = n.Val
	n.mu.Unlock()
	return val
}

func (n *Node[T]) LoadNext() *Node[T] {
	var next *Node[T]
	n.mu.Lock()
	next = n.Next
	n.mu.Unlock()
	return next
}

func (n *Node[T]) SwapVal(val T) T {
	n.mu.Lock()
	val, n.Val = n.Val, val
	n.mu.Unlock()
	return val
}

func (n *Node[T]) SwapNext(next *Node[T]) *Node[T] {
	n.mu.Lock()
	next, n.Next = n.Next, next
	n.mu.Unlock()
	return next
}
