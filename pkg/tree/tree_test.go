package tree

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func makeBinaryHelper[T any](xs []T) Binary[T] {
	root := Binary[T]{Val: xs[0]}
	a := []*Binary[T]{&root}
	l := len(xs)

	for i := range xs {
		node := a[i]

		leftChildIndex := 2*i + 1
		rightChildIndex := 2*i + 2

		if leftChildIndex < l {
			leftChild := Binary[T]{Val: xs[leftChildIndex]}
			node.Left = &leftChild
			a = append(a, &leftChild)
		}

		if rightChildIndex < l {
			rightChild := Binary[T]{Val: xs[rightChildIndex]}
			node.Right = &rightChild
			a = append(a, &rightChild)
		}
	}

	return *a[0]
}

func Test_InOrder(t *testing.T) {
	tree := makeBinaryHelper([]int{1, 2, 3, 4, 5})
	out := tree.InOrder()
	assert.Equal(t, []int{4, 2, 5, 1, 3}, out)
}
