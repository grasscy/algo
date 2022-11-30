package pass

var ans []int
var vis []bool

func grayCode(n int) []int {
	ans = []int{}
	vis = make([]bool, 1<<n)
	d(n, 0)
	return ans
}

func d(n int, cur int) {
	if len(ans) == 1<<n {
		return
	}
	ans = append(ans, cur)
	vis[cur] = true
	for i := 0; i < n; i++ {
		next := cur ^ (1 << i)
		if vis[next] {
			continue
		}
		d(n, next)
	}
	vis[cur] = false
}
