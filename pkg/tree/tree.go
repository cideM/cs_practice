package tree

type Binary[T comparable] struct {
	Val   T
	Left  *Binary[T]
	Right *Binary[T]
}

func (b Binary[T]) LNR() []T {
	out := []T{}

	if b.Left != nil {
		out = append(out, b.Left.LNR()...)
	}

	out = append(out, b.Val)

	if b.Right != nil {
		out = append(out, b.Right.LNR()...)
	}

	return out
}

func (b Binary[T]) NLR() []T {
	out := []T{}

	out = append(out, b.Val)

	if b.Left != nil {
		out = append(out, b.Left.NLR()...)
	}

	if b.Right != nil {
		out = append(out, b.Right.NLR()...)
	}

	return out
}

func (b Binary[T]) NRL() []T {
	out := []T{}

	out = append(out, b.Val)

	if b.Right != nil {
		out = append(out, b.Right.NRL()...)
	}

	if b.Left != nil {
		out = append(out, b.Left.NRL()...)
	}

	return out
}

func isMirror[T comparable](l, r *Binary[T]) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}

	return l.Val == r.Val && isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
}

func (b Binary[T]) IsSymmetric() bool {
	return isMirror(b.Left, b.Right)
}
