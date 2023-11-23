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

func (n *Node[T]) SetNext(next *Node[T]) {
	n.Mu.Lock()
	n.Next = next
	n.Mu.Unlock()
}

func (n *Node[T]) LoadVal() T {
	var val T
	n.Mu.Lock()
	val = n.Val
	n.Mu.Unlock()
	return val
}

func (n *Node[T]) LoadNext() *Node[T] {
	var next *Node[T]
	n.Mu.Lock()
	next = n.Next
	n.Mu.Unlock()
	return next
}

func (n *Node[T]) SwapVal(val T) T {
	n.Mu.Lock()
	val, n.Val = n.Val, val
	n.Mu.Unlock()
	return val
}

func (n *Node[T]) SwapNext(next *Node[T]) *Node[T] {
	n.Mu.Lock()
	next, n.Next = n.Next, next
	n.Mu.Unlock()
	return next
}
