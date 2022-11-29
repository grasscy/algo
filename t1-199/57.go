package pass

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append([][]int{}, newInterval)
	}
	ans := [][]int{}
	i := 0
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		ans = append(ans, intervals[i])
	}

	l := newInterval[0]
	r := newInterval[1]
	for ; i < len(intervals) && intervals[i][0] <= r; i++ {
		l = min(intervals[i][0], l)
		r = max(intervals[i][1], r)
	}

	ans = append(ans, []int{l, r})
	ans = append(ans, intervals[i:]...)

	return ans
}
