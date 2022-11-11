package ac

import (
	"fmt"
	"strconv"
)

var ks = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
var ans []string

func letterCombinations(digits string) []string {
	ans = []string{}
	if len(digits) == 0 {
		return ans
	}
	d(digits, 0, "")
	return ans
}

func d(digits string, index int, tmp string) {
	if len(tmp) == len(digits) {
		ans = append(ans, tmp)
		return
	}
	for i := index; i < len(digits); i++ {
		s := ks[toInt(digits[i])]
		for j := 0; j < len(s); j++ {
			d(digits, i+1, tmp+string(s[j]))
		}
	}
}

func toInt(u uint8) int {
	v, _ := strconv.Atoi(fmt.Sprintf("%c", u))
	return v
}
