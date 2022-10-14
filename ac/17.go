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
	b(digits, "", 0)
	return ans
}

func b(digits string, tmp string, index int) {
	if len(tmp) == len(digits) {
		ans = append(ans, tmp)
		return
	}
	k := ks[toInt(digits[index])]
	for i := 0; i < len(k); i++ {
		tmp += string(k[i])
		b(digits, tmp, index+1)
		tmp = tmp[:len(tmp)-1]
	}
}

func toInt(u uint8) int {
	v, _ := strconv.Atoi(fmt.Sprintf("%c", u))
	return v
}
