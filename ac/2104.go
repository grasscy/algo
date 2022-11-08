package ac

func subArrayRanges(nums []int) int64 {
	n := len(nums)
	mm := initArray(n, n)
	mn := initArray(n, n)
	for i := 0; i < n; i++ {
		mm[i][i] = nums[i]
		mn[i][i] = nums[i]
	}
	ans := 0
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			mm[i][j] = max(mm[i+1][j], nums[i])
			mn[i][j] = min(mn[i+1][j], nums[i])
			ans += mm[i][j] - mn[i][j]
		}
	}
	return int64(ans)
}
