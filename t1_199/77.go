package t1_199

var ans [][]int

//var vis []bool

func combine(n int, k int) [][]int {
	ans = [][]int{}
	b(n, k, 1, []int{})
	return ans
}

func b(n int, k int, start int, cur []int) {
	if len(cur) == k {
		ans = append(ans, append([]int{}, cur...))
		return
	}
	for i := start; i <= n; i++ {
		cur = append(cur, i)
		b(n, k, i+1, cur)
		cur = cur[:len(cur)-1]
	}
}
