package set

import "sync"

type Set[T comparable] struct {
	Mu   *sync.Mutex
	Data map[T]struct{}
}
