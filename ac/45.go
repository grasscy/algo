package ac

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	step := 0
	start := 0
	for {
		end := start + nums[start]
		if end >= len(nums)-1 {
			return step + 1
		}
		endi := start
		for i := 1; i <= nums[start]; i++ {
			if start+i+nums[start+i] > end {
				end = start + i + nums[start+i]
				endi = start + i
			}
		}
		start = endi
		step++
	}
	return step
}
