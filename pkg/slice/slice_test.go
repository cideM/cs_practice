package slice

import (
	"fmt"
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

func Test_RemoveDuplicates(t *testing.T) {
	input := []int{1, 1, 2}
	actual := RemoveDuplicates(input)
	assert.Equal(t, 2, actual)
	assert.Equal(t, []int{1, 2}, input[:actual])

	input = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	actual = RemoveDuplicates(input)
	assert.Equal(t, 5, actual)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, input[:actual])
}

func Test_BubbleSort(t *testing.T) {
	input := []int{1, 4, 2}
	BubbleSort(input)
	expected := []int{1, 2, 4}
	assert.Equal(t, expected, input)
}

func Test_MaxSubArray(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{in: []int{1}, expected: 1},
		{in: []int{-5, -5}, expected: -5},
		{in: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expected: 6},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d, %d", test.in, test.expected)
		t.Run(name, func(t *testing.T) {
			actual := MaxSubArray(test.in)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_PlusOne(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{in: []int{1}, expected: []int{2}},
		{in: []int{9}, expected: []int{1, 0}},
		{in: []int{1, 9}, expected: []int{2, 0}},
		{in: []int{1, 9, 9}, expected: []int{2, 0, 0}},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d, %d", test.in, test.expected)
		t.Run(name, func(t *testing.T) {
			actual := PlusOne(test.in)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_MergeSorged(t *testing.T) {
	tests := []struct {
		nums1, nums2, expected []int
		m, n                   int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			nums2:    []int{2, 5, 6},
			m:        3,
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1},
			nums2:    []int{},
			m:        1,
			n:        0,
			expected: []int{1},
		},
		{
			nums1:    []int{2, 0},
			nums2:    []int{1},
			m:        1,
			n:        1,
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v, %d, %v, %d, %v", test.nums1, test.m, test.nums2, test.n, test.expected)
		t.Run(name, func(t *testing.T) {
			MergeSorted(test.nums1, test.m, test.nums2, test.n)
			assert.Equal(t, test.expected, test.nums1)
		})
	}
}

func Test_MaxProfitIndeces(t *testing.T) {
	tests := []struct {
		in        []int
		expected1 int
		expected2 int
	}{
		{in: []int{}, expected1: -1, expected2: -1},
		{in: []int{7, 1, 5, 3, 6, 4}, expected1: 1, expected2: 4},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d, %d, %d", test.in, test.expected1, test.expected2)
		t.Run(name, func(t *testing.T) {
			a, b := MaxProfitIndeces(test.in)
			assert.Equal(t, test.expected1, a)
			assert.Equal(t, test.expected2, b)
		})
	}
}
