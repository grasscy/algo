package t200_399

var ans [][]int
var vis []bool

func combinationSum3(k int, n int) [][]int {
	ans = [][]int{}
	vis = make([]bool, 10)
	backtrack(k, n, 0, 0, []int{}, 1)
	return ans
}

func backtrack(k int, n int, sum int, count int, cur []int, start int) {
	if sum == n && count == k {
		ans = append(ans, append([]int{}, cur...))
	}
	if sum > n || count > k {
		return
	}
	for i := start; i <= 9; i++ {
		if vis[i] {
			continue
		}
		vis[i] = true
		cur = append(cur, i)
		backtrack(k, n, sum+i, count+1, cur, i+1)
		cur = cur[:len(cur)-1]
		vis[i] = false
	}
}
