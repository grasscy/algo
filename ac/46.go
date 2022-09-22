package ac

var ans [][]int
var vis []bool

func permute(nums []int) [][]int {
	ans = [][]int{}
	vis = make([]bool, len(nums))
	b(nums, []int{})
	return ans
}

func b(nums []int, cur []int) {
	if len(nums) == len(cur) {
		ans = append(ans, append([]int{}, cur...))
	}
	for i := 0; i < len(nums); i++ {
		if vis[i] {
			continue
		}
		vis[i] = true
		cur = append(cur, nums[i])
		b(nums, cur)
		cur = cur[:len(cur)-1]
		vis[i] = false
	}
}
