package pass

func maxArea(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	for l < r {
		if height[l] < height[r] {
			if ans < (r-l)*height[l] {
				ans = (r - l) * height[l]
			}
			l++
		} else {
			if ans < (r-l)*height[r] {
				ans = (r - l) * height[r]
			}
			r--
		}
	}
	return ans
}
