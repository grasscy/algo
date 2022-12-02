package t1_199

var ans [][]int

func subsets(nums []int) [][]int {
	ans = [][]int{}
	b(nums, 0, []int{})
	return ans
}

func b(nums []int, start int, cur []int) {
	ans = append(ans, append([]int{}, cur...))

	for i := start; i < len(nums); i++ {
		cur = append(cur, nums[i])
		b(nums, i+1, cur)
		cur = cur[:len(cur)-1]
	}
}
