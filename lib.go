package main

import "math"

// LongestCommonPrefix uses vertical scanning, so it looks at one column of
// letters at a time.
// The outer loop has $shortest_str_len cycles, the inner loop $num_strs, so
// it's O($shortest * $num_strs). Worst case would be for all strings to have
// the same length
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	min := math.MaxInt
	for _, s := range strs {
		if len(s) < min {
			min = len(s)
		}
	}

	var offset int
	for i := 0; i < min; i++ {
		for _, s := range strs {
			if s[i] != strs[0][i] {
				return strs[0][:offset]
			}
		}
		offset = offset + 1
	}
	return strs[0][:offset]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeSortedLinkedLists looks at the current node of both lists and adds
// the smaller one to the new list. It then advances the list from which it
// took the node. Therefore, both lists are advanced separately, and
// traversed at max once. If the end of one list is reached, we can just
// append the remainder of the other list. The reason this all works is
// because both input lists are already sorted. This is O(n + m) with n and m being the lengths of the lists.
func MergeSortedLinkedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur1, cur2 := list1, list2
	root := &ListNode{}
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
