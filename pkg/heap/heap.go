package heap

// http://cs.middlesexcc.edu/~schatz/csc236/handouts/heapsort.html

import (
	"golang.org/x/exp/constraints"
	"math"
)

type compareFn[T constraints.Ordered] func(T, T) bool

type Heap[T constraints.Ordered] struct {
	values  []T
	compare func(T, T) bool
}

func New[T constraints.Ordered](values []T, fn compareFn[T]) Heap[T] {
	heapify(values, fn)
	return Heap[T]{values: values, compare: fn}
}

func (h *Heap[T]) Push(v T) {
	h.values = append(h.values, v)
	siftUp(h.values, len(h.values)-1, h.compare)
}

func (h *Heap[T]) Pop() {
  last := len(h.values) - 1
  h.values[0] = h.values[last]
  h.values = h.values[:last]
  heapify(h.values, h.compare)
}

func getLeftChildIndex(i int) int {
	return 2*i + 1
}

func getRightChildIndex(i int) int {
	return 2*i + 2
}

func getParentIndex(i int) int {
	return int(math.Floor(float64((i - 1) / 2)))
}

func heapify[T constraints.Ordered](xs []T, compare compareFn[T]) {
	l := len(xs)
	// leafs are valid heaps, so we can start at the last parent
	startI := getParentIndex(l - 1)

	for startI >= 0 {
		siftDown(xs, startI, l-1, compare)
		startI--
	}
}

func siftDown[T constraints.Ordered](xs []T, start, end int, compare compareFn[T]) {
	rootI := start

	for getLeftChildIndex(rootI) <= end {
		swapI := rootI
		leftChildI := getLeftChildIndex(rootI)

		if compare(xs[swapI], xs[leftChildI]) {
			swapI = leftChildI
		}

		if leftChildI+1 <= end && compare(xs[swapI], xs[leftChildI+1]) {
			swapI = leftChildI + 1
		}

		if swapI == rootI {
			return
		} else {
			xs[rootI], xs[swapI] = xs[swapI], xs[rootI]
			rootI = swapI
		}
	}
}

func siftUp[T constraints.Ordered](xs []T, start int, compare compareFn[T]) {
	nodeI := start

	for nodeI >= 1 {
		parentI := getParentIndex(nodeI)

		if !compare(xs[nodeI], xs[parentI]) {
			xs[nodeI], xs[parentI] = xs[parentI], xs[nodeI]
			nodeI = parentI
		} else {
			return
		}
	}
}
