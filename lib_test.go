package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestCommonPrefix(t *testing.T) {
	actual := LongestCommonPrefix([]string{
		"flower", "flow", "flight",
	})
	assert.Equal(t, "fl", actual)

	actual = LongestCommonPrefix([]string{})
	assert.Equal(t, "", actual)
}

func makeLinkedList(nums []int) *ListNode {
	root := &ListNode{}
	head := root
	for _, n := range nums {
		node := &ListNode{Val: n}
		root.Next = node
		root = node
	}
	return head.Next
}

func listToNums(n *ListNode) []int {
	cur := n
	nums := make([]int, 0)
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}
	return nums
}

func Test_MergeSortedLinkedLists(t *testing.T) {
	actual := listToNums(MergeSortedLinkedLists(
		makeLinkedList([]int{1, 2, 4}),
		makeLinkedList([]int{1, 3, 4}),
	))
	expected := listToNums(makeLinkedList([]int{1, 1, 2, 3, 4, 4}))
	assert.Equal(t, expected, actual)

	actual = listToNums(MergeSortedLinkedLists(
		makeLinkedList(nil),
		makeLinkedList([]int{1, 3, 4}),
	))
	expected = listToNums(makeLinkedList([]int{1, 3, 4}))
	assert.Equal(t, expected, actual)

	actual = listToNums(MergeSortedLinkedLists(nil, nil))
	expected = listToNums(makeLinkedList(nil))
	assert.Equal(t, expected, actual)
}
