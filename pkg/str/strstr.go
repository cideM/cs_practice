package search

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	match := 0
	l := len(needle)
	for pos, c := range haystack {
		if match == l {
			return pos - l - 1
		}
		if byte(c) == needle[match] {
			match++
			continue
		} else if byte(c) == needle[0] {
			match = 1
			continue
		}
		match = 0
	}

	return -1
}
