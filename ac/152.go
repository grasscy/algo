package ac

func maxProduct(nums []int) int {
	mmax := make([]int, len(nums))
	mmin := make([]int, len(nums))
	mmin[0] = nums[0]
	mmax[0] = nums[0]

	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		mmax[i] = max(max(mmax[i-1]*nums[i], nums[i]), mmin[i-1]*nums[i])
		mmin[i] = min(min(mmin[i-1]*nums[i], nums[i]), mmax[i-1]*nums[i])
		ans = max(mmax[i], ans)
	}
	return ans

}
