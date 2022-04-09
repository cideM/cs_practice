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
