package heap

import (
	"golang.org/x/exp/constraints"
)

type MaxHeap[T constraints.Ordered] struct {
  Heap[T]
}

func NewMax[T constraints.Ordered](values []T) MaxHeap[T] {
  fn := func(a T, b T) bool {
    return a < b
  }
  return MaxHeap[T]{Heap: New(values, fn)}
}

func (m MaxHeap[T]) Max() T {
	return m.Heap.values[0]
}

