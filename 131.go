package algo

import "fmt"

var ans [][]string
var cs []string

func partition(s string) [][]string {
	ans = [][]string{}
	cs = []string{}
	b(s, "", 0)
	return ans
}

func b(s string, tmp string, index int) {
	cs = append(cs, tmp)
	fmt.Println(cs)
	for i := index; i < len(s); i++ {
		tmp += string(s[i])
		b(s, tmp, i+1)
		tmp = tmp[:len(tmp)-1]
	}

}

//
//func check(s string) bool {
//
//}
