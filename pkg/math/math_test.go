package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SquareRootBinarySearch(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{in: 1, expected: 1},
		{in: 4, expected: 2},
		{in: 11, expected: 3},
		{in: 2147395600, expected: 46340},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d, %d", test.in, test.expected)
		t.Run(name, func(t *testing.T) {
			actual := SquareRootBinarySearch(test.in)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_SquareRootNewton(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{in: 1, expected: 1},
		{in: 4, expected: 2},
		{in: 8, expected: 2},
		{in: 11, expected: 3},
		{in: 2147395600, expected: 46340},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d, %d", test.in, test.expected)
		t.Run(name, func(t *testing.T) {
			actual := SquareRootNewton(test.in)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_PascalTriangle(t *testing.T) {
	tests := []struct {
		in       int
		expected [][]int
	}{
		{in: 0, expected: [][]int{}},
		{in: 1, expected: [][]int{{1}}},
		{in: 2, expected: [][]int{{1},{1,1}}},
		{in: 5, expected: [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d, %v", test.in, test.expected)
		t.Run(name, func(t *testing.T) {
			actual := PascalTriangle(test.in)
			assert.Equal(t, test.expected, actual)
		})
	}
}
