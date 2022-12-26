package t1_199

func findMin(nums []int) int {
	ans := nums[0]
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		ans = min(ans, nums[l])
		ans = min(ans, nums[mid])
		if nums[l] < nums[mid] {
			l = mid + 1
		} else if nums[l] > nums[mid] {
			r = mid - 1
		} else {
			l++
		}
	}
	return ans
}
