package slice

import (
	"golang.org/x/exp/constraints"
	"math"
)

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

// https://go.dev/ref/spec#General_interfaces
// > Implementation restriction: A union (with more than one term) cannot
// > contain the predeclared identifier comparable or interfaces that specify
// > methods, or embed comparable or interfaces that specify methods. 
type orderedComparable interface {
  constraints.Ordered
  comparable
}

// https://en.wikipedia.org/wiki/Sorting_algorithm

// BubbleSort is simple but slow. It goes over the list repeatedly until
// everything is sorted. It's stable and has O(n) memory.
func BubbleSort[T orderedComparable](xs []T) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(xs); i++ {
			if xs[i-1] > xs[i] {
				swapped = true
				xs[i-1], xs[i] = xs[i], xs[i-1]
			}
		}
	}
}


// RemoveDuplicates only works on sorted slices and returns the number of
// elements of the input slice that are sorted. It's O(n) where n is the length
// of the slice.
func RemoveDuplicates[T orderedComparable](xs []T) int {
	l := len(xs)
	if l < 2 {
		return l
	}
	pos := 0
	for i := 1; i < l; i++ {
		if xs[pos] != xs[i] {
			pos++
			xs[pos] = xs[i]
		}
	}
	return pos + 1
}
