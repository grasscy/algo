package pass

import "math"

func minWindow(s string, t string) string {
	need := map[string]int{}
	for _, v := range t {
		need[string(v)]++
	}
	left, right := 0, 0
	valid := 0
	window := map[string]int{}
	ans := math.MaxInt64
	start := 0
	for ; right < len(s); right++ {
		c := string(s[right])

		window[c]++
		if need[c] == window[c] {
			valid++
		}

		for valid == len(need) {
			if right-left < ans {
				ans = right - left
				start = left
			}

			c := string(s[left])
			left++

			if _, ok := need[c]; ok {
				if window[c] == need[c] {
					valid--
				}
				window[c]--
			}
		}
	}
	if ans == math.MaxInt64 {
		return ""
	}
	return s[start : start+ans+1]
}
