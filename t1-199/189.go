package pass

func rotate(nums []int, k int) {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = ans[i]
	}
	return
}
