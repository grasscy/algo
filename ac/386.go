package ac

var ans []int

func lexicalOrder(n int) []int {
	ans = []int{}
	dfs(n, 0)
	return ans
}

func dfs(n int, index int) {
	for i := 0; i <= 9; i++ {
		if index == 0 && i == 0 {
			continue
		}
		t := index*10 + i
		if t > n {
			return
		}
		ans = append(ans, t)
		dfs(n, t)
	}
}
