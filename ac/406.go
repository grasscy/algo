package ac

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})

	r := [][]int{}

	for _, v := range people {
		if len(r) <= v[1] {
			r = append(r, v)
		} else {
			pre := append([][]int{}, r[:v[1]]...)
			suf := append([][]int{}, r[v[1]:]...)

			r = append(append(pre, v), suf...)
		}
	}
	return r
}
