package ac

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	ans := [][]int{}
	for i := 0; i < len(people); i++ {
		v := people[i]
		if v[1] >= len(ans) {
			ans = append(ans, v)
		} else {
			pre := append([][]int{}, ans[:v[1]]...)
			suf := append([][]int{}, ans[v[1]:]...)

			ans = append(append(pre, v), suf...)
		}
	}
	return ans
}
