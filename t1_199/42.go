package t1_199

func trap(height []int) int {
	ans := 0
	lefts := make([]int, len(height))
	rights := make([]int, len(height))
	for i := 1; i < len(height)-1; i++ {
		lefts[i] = max(lefts[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rights[i] = max(rights[i+1], height[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		mmin := min(lefts[i], rights[i])
		if mmin > height[i] {
			ans += mmin - height[i]
		}
	}
	return ans
}
