// Package singly provides generic Singly Linked Node (SLN) data structure and its basic methods.
package singly

import "sync"

// Singly Linked Node (SLN) is a data structure that can store a value and can be linked to the next Node.
//
// SLN consists of value of type T and a pointer to the next Node.
//
// Thread-safe approach is to use implemented methods.
type Node[T any] struct {
	mu    *sync.Mutex
	Value T
	Next  *Node[T]
}

// New creates a Node of type T.
func New[T any]() *Node[T] {
	return &Node[T]{
		mu: &sync.Mutex{},
	}
}

// SetValue stores value into the Node.
func (n *Node[T]) SetValue(value T) {
	n.mu.Lock()
	n.Value = value
	n.mu.Unlock()
}

// SetValue overwrites "Next" pointer in the Node.
func (n *Node[T]) SetNext(next *Node[T]) {
	n.mu.Lock()
	n.Next = next
	n.mu.Unlock()
}

// LoadValue returns value from the Node.
func (n *Node[T]) LoadValue() T {
	var value T
	n.mu.Lock()
	value = n.Value
	n.mu.Unlock()
	return value
}

// LoadNext returns "Next" pointer from the Node.
func (n *Node[T]) LoadNext() *Node[T] {
	var next *Node[T]
	n.mu.Lock()
	next = n.Next
	n.mu.Unlock()
	return next
}

// SwapValue stores value into Node and returns its previous value.
func (n *Node[T]) SwapValue(value T) T {
	n.mu.Lock()
	value, n.Value = n.Value, value
	n.mu.Unlock()
	return value
}

// SwapNext overwrites "Next" pointer in the Node and returns previous "Next" pointer.
func (n *Node[T]) SwapNext(next *Node[T]) *Node[T] {
	n.mu.Lock()
	next, n.Next = n.Next, next
	n.mu.Unlock()
	return next
}
