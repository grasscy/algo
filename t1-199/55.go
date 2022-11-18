package pass

func canJump(nums []int) bool {
	maxi := nums[0]

	for i := 1; i < len(nums); i++ {
		if maxi < i {
			return false
		}
		if maxi < i+nums[i] {
			maxi = i + nums[i]
		}
	}

	if maxi >= len(nums)-1 {
		return true
	}
	return false

}
