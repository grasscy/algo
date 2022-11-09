package ac

var vis []bool
var ans int

func countArrangement(n int) int {
	vis = make([]bool, n+1)
	ans = 0
	bt(n, []int{})
	return ans
}

func bt(n int, tmp []int) {
	if len(tmp) == n {
		ans++
	}
	for i := 1; i <= n; i++ {
		if vis[i] {
			continue
		}
		if !(i%(len(tmp)+1) == 0 || (len(tmp)+1)%i == 0) {
			continue
		}
		vis[i] = true
		bt(n, append(tmp, i))
		vis[i] = false
	}

}
