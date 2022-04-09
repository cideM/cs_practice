package linkedlist

import (
	"golang.org/x/exp/constraints"
)

type ListNode[T constraints.Ordered] struct {
	Val  T
	Next *ListNode[T]
}

// MergeSorted looks at the current node of both lists and adds the smaller one
// to the new list. It then advances the list from which it took the node.
// Therefore, both lists are advanced separately, and traversed at max once. If
// the end of one list is reached, we can just append the remainder of the
// other list. The reason this all works is because both input lists are
// already sorted. This is O(n + m) with n and m being the lengths of the
// lists.
func MergeSorted[T constraints.Ordered](list1 *ListNode[T], list2 *ListNode[T]) *ListNode[T] {
	cur1, cur2 := list1, list2
	root := &ListNode[T]{}
	head := root

	for {
		if cur1 == nil {
			root.Next = cur2
			return head.Next
		}

		if cur2 == nil {
			root.Next = cur1
			return head.Next
		}

		if cur1.Val < cur2.Val {
			p := cur1
			root.Next = p
			root = p
			cur1 = cur1.Next
			continue
		}

		p := cur2
		root.Next = p
		root = p
		cur2 = cur2.Next

	}
}
