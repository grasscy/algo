package ac

var ans int

func findTargetSumWays(nums []int, target int) int {
	ans = 0
	d(nums, target, 0, 0)
	return ans
}

func d(nums []int, tg int, sum int, index int) {
	if index == len(nums) {
		if tg == sum {
			ans++
		}
		return
	}
	i := index
	d(nums, tg, sum-nums[i], i+1)
	d(nums, tg, sum+nums[i], i+1)
}
