package math

import (
	"math"
)

// SquareRootBinarySearch finds the square root of "x" using binary search to
// get a better approximation on each iteration
func SquareRootBinarySearch(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	start, end, guess := 0, x/2, 0
	for start <= end {
		mid := (start + end) / 2
		square := mid * mid

		if square == x {
			return mid
		}

		if square > x {
			end--
			continue
		} else if square < x {
			guess = mid
			start++
			continue
		}
	}

	// this returns the best guess, e.g. 11 doesn't have a clean square root
	return guess
}

// SquareRootNewton uses https://math.mit.edu/~stevenj/18.335/newton-sqrt.pdf
// This actually implements the original formula instead of the copy & pasted
// variant that every tutorial on the interweb implements because people just
// copy from each other
func SquareRootNewton(n int) int {
	var LIMIT = 0.00001
	if n == 0 || n == 1 {
		return n
	}

	a := float64(n)
	guess := a / 2

	for {
		newGuess := 0.5 * float64(guess+a/guess)

		if math.Abs(newGuess-guess) <= LIMIT {
			return int(guess)
		}
		guess = newGuess
	}
}
