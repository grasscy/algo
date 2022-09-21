package ac

func jump(nums []int) int {
	start := 0
	end := 0
	ans := 0
	for end < len(nums)-1 {
		mp := 0
		for i := start; i <= end; i++ {
			mp = max(mp, i+nums[i])
		}
		start = end + 1
		end = mp
		ans++
	}

	return ans
}
