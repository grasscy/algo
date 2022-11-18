package pass

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	s, e := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > e {
			ans = append(ans, []int{s, e})
			s = intervals[i][0]
			e = intervals[i][1]
		} else {
			if e < intervals[i][1] {
				e = intervals[i][1]
			}
		}
	}
	ans = append(ans, []int{s, e})
	return ans
}
