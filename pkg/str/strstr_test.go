package str

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_strStr(t *testing.T) {
	tests := []struct {
		haystack, needle string
		expected         int
	}{
		{haystack: "foobar", needle: "ob", expected: 2},
		{haystack: "foobar", needle: "", expected: 0},
		{haystack: "foobar", needle: "f", expected: 0},
		{haystack: "foobar", needle: "foobar", expected: 0},
		{haystack: "foob ar", needle: "ar", expected: 5},
		{haystack: "babba", needle: "bbb", expected: -1},
		{haystack: "mississippi", needle: "issip", expected: 4},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%s %s %d", test.haystack, test.needle, test.expected)
		t.Run(name, func(t *testing.T) {
			actual := strStr(test.haystack, test.needle)
			assert.Equal(t, test.expected, actual)
		})
	}
}
