package algo

func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	intervals = append([][]int{{-2, -1}}, intervals...)

	i := 0
	for ; i < len(intervals)-1 && intervals[i+1][0] < newInterval[0]; i++ {
	}
	ans = append(ans, intervals[1:i]...)
	j := i
	for ; j < len(intervals)-1 && intervals[j+1][1] < newInterval[1]; j++ {
	}

	return ans
}
