package str

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
