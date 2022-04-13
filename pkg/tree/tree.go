package tree

type Binary[T any] struct {
	Val   T
	Left  *Binary[T]
	Right *Binary[T]
}

// Left, Node, Right
func (b Binary[T]) InOrder() []T {
	out := []T{}

	if b.Left != nil {
		out = append(out, b.Left.InOrder()...)
	}

	out = append(out, b.Val)

	if b.Right != nil {
		out = append(out, b.Right.InOrder()...)
	}

	return out
}
