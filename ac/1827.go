package ac

func minOperations(nums []int) int {
	ans := 0
	if len(nums) == 1 {
		return ans
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			ans += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return ans
}
