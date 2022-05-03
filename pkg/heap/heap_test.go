package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Min(t *testing.T) {
	input := []int{6, 5, 3, 1, 8, 7, 2, 4}
	h := NewMin(input)
	assert.Equal(t, 1, h.Min())

	h.Push(0)
	assert.Equal(t, 0, h.Min())

  h.Pop()
	assert.Equal(t, 1, h.Min())
}

func Test_Max(t *testing.T) {
	input := []int{6, 5, 3, 1, 8, 7, 2, 4}
	h := NewMax(input)
	assert.Equal(t, 8, h.Max())

	h.Push(10)
	assert.Equal(t, 10, h.Max())

  h.Pop()
	assert.Equal(t, 8, h.Max())
}
