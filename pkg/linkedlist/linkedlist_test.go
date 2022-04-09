package linkedlist

import (
	"testing"

	"golang.org/x/exp/constraints"
	"github.com/stretchr/testify/assert"
)

func makeLinkedList[T constraints.Ordered](nums []T) *ListNode[T] {
	root := &ListNode[T]{}
	head := root
	for _, n := range nums {
		node := &ListNode[T]{Val: n}
		root.Next = node
		root = node
	}
	return head.Next
}

func unroll[T constraints.Ordered](n *ListNode[T]) []T {
	cur := n
	nums := make([]T, 0)
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}
	return nums
}

func Test_MergeSorted(t *testing.T) {
	actual := unroll(MergeSorted(
		makeLinkedList([]int{1, 2, 4}),
		makeLinkedList([]int{1, 3, 4}),
	))
	expected := unroll(makeLinkedList([]int{1, 1, 2, 3, 4, 4}))
	assert.Equal(t, expected, actual)

	actual = unroll(MergeSorted(
		makeLinkedList[int](nil),
		makeLinkedList([]int{1, 3, 4}),
	))
	expected = unroll(makeLinkedList([]int{1, 3, 4}))
	assert.Equal(t, expected, actual)

	actual = unroll(MergeSorted[int](nil, nil))
	expected = unroll(makeLinkedList[int](nil))
	assert.Equal(t, expected, actual)
}
