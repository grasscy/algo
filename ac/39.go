package ac

var ans [][]int

func combinationSum(candidates []int, target int) [][]int {
	ans = [][]int{}
	b(candidates, target, 0, []int{}, 0)
	return ans
}

func b(candidates []int, target int, sum int, cur []int, start int) {
	if sum == target {
		ans = append(ans, append([]int{}, cur...))
	}
	if sum > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		b(candidates, target, sum+candidates[i], cur, i)
		cur = cur[:len(cur)-1]
	}
}
