package t1_199

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	p := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[p-1] && nums[p-1] == nums[p-2] {
		} else {
			nums[p] = nums[i]
			p++
		}
	}
	return p

}
