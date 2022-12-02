package t1_199

import (
	"fmt"
	"strconv"
)

var ans []string

func restoreIpAddresses(s string) []string {
	ans = []string{}
	b(s, []int{}, 0)
	return ans
}

func b(s string, tmp []int, index int) {
	if len(tmp) == 4 && index == len(s) {
		ans = append(ans, fmt.Sprintf("%d.%d.%d.%d", tmp[0], tmp[1], tmp[2], tmp[3]))
	}

	if len(tmp) > 4 || index >= len(s) {
		return
	}

	if s[index] == '0' {
		b(s, append(tmp, 0), index+1)
		return
	}

	for i := index; i < index+3 && i < len(s); i++ {
		atoi, _ := strconv.Atoi(s[index : i+1])
		if atoi > 255 {
			continue
		}
		b(s, append(tmp, atoi), i+1)
	}
}
