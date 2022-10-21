package ac

var ans bool
var vis []bool

func makesquare(matchsticks []int) bool {
	ans = false
	vis = make([]bool, len(matchsticks))
	sum := 0
	for _, v := range matchsticks {
		sum += v
	}
	if sum%4 != 0 {
		return false
	}
	d(matchsticks, 0, 0, sum/4, 0)
	return ans
}

func d(matchsticks []int, n int, sum int, target int, index int) {
	if n == 4 {
		ans = true
	}
	if ans {
		return
	}
	if sum == target {
		d(matchsticks, n+1, 0, target, 0)
	}

	if sum > target {
		return
	}

	for i := index; i < len(matchsticks); i++ {
		if vis[i] {
			continue
		}
		vis[i] = true
		d(matchsticks, n, sum+matchsticks[i], target, i)
		vis[i] = false
	}

}
