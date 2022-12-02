package t1_199

var vis []bool
var ans [][]int

func permute(nums []int) [][]int {
	vis = make([]bool, len(nums))
	ans = [][]int{}
	b(nums, []int{})
	return ans
}

func b(nums []int, tmp []int) {
	if len(tmp) == len(nums) {
		ans = append(ans, append([]int{}, tmp...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if vis[i] {
			continue
		}
		vis[i] = true
		b(nums, append(tmp, nums[i]))
		vis[i] = false
	}
}
