package t1_199

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		k := i
		j := 0
		for k < len(haystack) && j < len(needle) && haystack[k] == needle[j] {
			k++
			j++
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
