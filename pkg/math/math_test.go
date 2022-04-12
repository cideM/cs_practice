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
