package ac

import (
	"fmt"
	"strconv"
)

var r []string
var ks = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	r = []string{}
	if len(digits) == 0 {
		return r
	}

	f(digits, 0, "")

	return r
}

func f(digits string, index int, s string) {
	if index == len(digits) {
		r = append(r, s)
		return
	}
	ls := ks[toInt(digits[index])]
	for j := 0; j < len(ls); j++ {
		f(digits, index+1, s+string(ls[j]))
	}

}

func toInt(u uint8) int{
	v, _:=strconv.Atoi(fmt.Sprintf("%c", u))
	return v
}