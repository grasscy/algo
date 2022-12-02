package t1_199

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1

}
