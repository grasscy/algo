package t1_199

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	ns := []string{}
	for _, v := range nums {
		ns = append(ns, strconv.Itoa(v))
	}
	sort.Slice(ns, func(i, j int) bool {
		s1 := ns[i]
		s2 := ns[j]

		k := 0
		for ; k < len(s1) && k < len(s2); k++ {
			if s1[k] > s2[k] {
				return true
			} else if s1[k] < s2[k] {
				return false
			}
		}
		if s1+s2 > s2+s1 {
			return true
		}
		return false
	})

	var i int
	for ; i < len(ns) && ns[i] == "0"; i++ {

	}
	if i == len(ns) {
		return "0"
	}
	join := strings.Join(ns[i:], "")

	return join
}
