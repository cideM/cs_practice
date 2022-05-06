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

func (n ListNode[T]) HasCycle() bool {
	// Brent algorithm
	// https://en.wikipedia.org/wiki/Cycle_detection#Brent's_algorithm
  if n.Next == nil {
    return false
  }
	turtle, rabbit := n.Next, n.Next

	step_limit := 2
	steps_taken := 0

	for {
		if rabbit.Next == nil {
			return false
		}

		// Advance the rabbit until we've reached the steps_taken limit, which we
		// keep increasing after each sequence of steps. We increase it by a power
		// of two. That way the rabbit runs ahead and the turtle is occasionalyl
		// teleported to the rabbit.
		rabbit = rabbit.Next

		steps_taken++

		if rabbit == turtle {
			return true
		}

		if steps_taken == step_limit {
			steps_taken = 0
			step_limit *= 2
			turtle = rabbit
		}
	}
}

func (n ListNode[T]) DoesIntersect(other ListNode[T]) *ListNode[T] {
  a, b := &n, &other

  for a != b {
    if a == nil && b == nil {
      return nil
    }

    if a == nil {
      a = &other
    } else {
      a = a.Next
    }

    if b == nil {
      b = &n
    } else {
      b = b.Next
    }
  }

  return a
}
