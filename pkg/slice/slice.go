package slice

import (
	"golang.org/x/exp/constraints"
	"math"
	// "log"
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

// MaxSubArray uses Kadane's algorithm and expects an input list of at least
// length 1
func MaxSubArray(nums []int) int {
	cur, best := 0, math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum := cur + nums[i]
		if sum > nums[i] {
			cur = sum
		} else {
			cur = nums[i]
		}
		if cur > best {
			best = cur
		}
	}
	return best
}

// PlusOne treats a slice of digits like a number and increments it by one.
// Numbers must be greater than or equal to zero.
func PlusOne(digits []int) []int {
	l := len(digits)

	carry := true
	for i := l - 1; i >= 0; i-- {
		if !carry {
			return digits
		}

		cur := digits[i]

		if i == 0 {
			if cur == 9 {
				newInts := []int{1, 0}
				newInts = append(newInts, digits[1:]...)
				return newInts
			}

			digits[i]++
			return digits
		}

		if cur == 9 {
			digits[i] = 0
			continue
		}

		digits[i]++
		carry = false
	}

	return digits
}

// MergeSorted merges nums2 into nums1 while maintaining the sorting. nums1
// will have enough space and the length of the actual array num1 is denoted by
// m. Both nums1 nums2 are sorted.
func MergeSorted(nums1 []int, m int, nums2 []int, n int) {
	// This is the second attempt after reading some discussions online. It's
	// much better to start at the end of both slices and then also start
	// **writing to the end of nums1**. This last part is the key insight,
	// because then you don't need any swapping.
	// Think about it: with the way the question is phrased we know that we will
	// fill nums1 eventually. So whichever last element from nums1 and nums2 is
	// greater, it will end up in the last position of nums1.
	write := m + n - 1
	m--
	n--
	for n >= 0 {
		if m >= 0 && nums1[m] > nums2[n] {
			nums1[write] = nums1[m]
			m--
		} else {
			nums1[write] = nums2[n]
			n--
		}
		write--
	}
}

// MergeSortedFirst was my first attempt which starts at the beginning of both
// slices. The problem is that it will do a lot of shifting and as such it'll
// be slower than it has to.
func MergeSortedFirst(nums1 []int, m int, nums2 []int, n int) {
	i1, i2 := 0, 0
	for i2 < n {
		if i1 >= m {
			copy(nums1[m:], nums2[i2:])
			return
		}

		for i2 < n && nums2[i2] < nums1[i1] {
			copy(nums1[i1+1:], nums1[i1:])
			nums1[i1] = nums2[i2]
			m++
			i2++
			continue
		}
		i1++
	}
}
