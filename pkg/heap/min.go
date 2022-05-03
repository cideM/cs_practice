package heap

import (
	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] struct {
  Heap[T]
}

func NewMin[T constraints.Ordered](values []T) MinHeap[T] {
  fn := func(a T, b T) bool {
    return a > b
  }
  return MinHeap[T]{Heap: New(values, fn)}
}

func (m MinHeap[T]) Min() T {
	return m.Heap.values[0]
}


