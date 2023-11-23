package singly

import "sync"

type Node[T any] struct {
	Mu   *sync.Mutex
	Val  T
	Next *Node[T]
}
