package ac

func findUnsortedSubarray(nums []int) int {
	max, min := nums[0], nums[len(nums)-1]
	start, end := 0, -1

	for i := 0; i < len(nums); i++ {
		if nums[i] < max {
			end = i
		} else {
			max = nums[i]
		}

		if nums[len(nums)-1-i] > min {
			start = len(nums) - 1 - i
		} else {
			min = nums[len(nums)-1-i]
		}
	}
	return end - start + 1

}
