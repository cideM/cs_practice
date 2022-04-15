package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func makeBinaryHelper[T comparable](xs []*T) *Binary[T] {
  if xs[0] == nil {
    return nil
  }

	root := Binary[T]{Val: *xs[0]}
	a := []*Binary[T]{&root}
	l := len(xs)

	for i := range xs {
		node := a[i]

		leftChildIndex := 2*i + 1
		rightChildIndex := 2*i + 2

		if leftChildIndex < l {
			leftV := xs[leftChildIndex]
			var leftChild *Binary[T]
			if leftV != nil {
				leftChild = &Binary[T]{Val: *leftV}
			}
			node.Left = leftChild
			a = append(a, leftChild)
		}

		if rightChildIndex < l {
			rightV := xs[rightChildIndex]
			var rightChild *Binary[T]
			if rightV != nil {
				rightChild = &Binary[T]{Val: *rightV}
			}
			node.Right = rightChild
			a = append(a, rightChild)
		}
	}

	return a[0]
}

func intp(i int) *int {
	return &i
}

func Test_LNR(t *testing.T) {
	tree := makeBinaryHelper([]*int{intp(1), intp(2), intp(3), intp(4), intp(5)})
	out := tree.LNR()
	assert.Equal(t, []int{4, 2, 5, 1, 3}, out)
}

func Test_IsSymmetric(t *testing.T) {
	tree := makeBinaryHelper([]*int{intp(1), intp(2), intp(2), intp(3), intp(4), intp(4), intp(3)})
	assert.True(t, tree.IsSymmetric())

	tree2 := makeBinaryHelper([]*int{intp(1), intp(2), intp(2), intp(2), nil, intp(2)})
	assert.False(t, tree2.IsSymmetric())

	tree3 := makeBinaryHelper([]*int{intp(1), intp(2), intp(2), nil, intp(3), nil, intp(3)})
	assert.False(t, tree3.IsSymmetric())
}

func Test_MaxDepth(t *testing.T) {
	tree := makeBinaryHelper([]*int{intp(1), intp(2), intp(3), intp(4), intp(5)})
	out := tree.MaxDepth()
	assert.Equal(t, 3, out)
}

