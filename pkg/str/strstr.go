package str

import (
	"unicode"
	"unicode/utf8"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	l := len(needle)
	last := len(haystack)
	for pos := range haystack {
		if last-pos < l {
			return -1
		}
		if haystack[pos:pos+l] == needle {
			return pos
		}
	}

	return -1
}

// palindrome is case insensitive and only considers letters and numbers but
// works with unicode
func palindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		lr, ls := utf8.DecodeRuneInString(s[l:])
		if lr == utf8.RuneError || (!unicode.IsLetter(lr) && !unicode.IsNumber(lr)) {
			l++
			continue
		}

		rr, rs := utf8.DecodeRuneInString(s[r:])
		if rr == utf8.RuneError || (!unicode.IsLetter(rr) && !unicode.IsNumber(rr)) {
			r--
			continue
		}

		if unicode.ToLower(lr) != unicode.ToLower(rr) {
			return false
		} else {
			l += ls
			r -= rs
		}
	}

	return true
}
