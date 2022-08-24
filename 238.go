package algo

func productExceptSelf(nums []int) []int {
	r := make([]int, len(nums))

	r[0] = 1
	for i := 1; i < len(nums); i++ {
		r[i] = r[i-1] * nums[i-1]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		r[i] = r[i] * right
		right *= nums[i]
	}
	return r
}
