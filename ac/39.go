package ac

var ans [][]int

func combinationSum(candidates []int, target int) [][]int {
	ans = [][]int{}
	b(candidates, target, []int{}, 0, 0)
	return ans
}

func b(candidates []int, target int, tmp []int, sum int, index int) {
	if sum == target {
		ans = append(ans, append([]int{}, tmp...))
		return
	}
	if sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		sum += candidates[i]
		tmp = append(tmp, candidates[i])
		b(candidates, target, tmp, sum, i)
		tmp = tmp[:len(tmp)-1]
		sum -= candidates[i]
	}
}
