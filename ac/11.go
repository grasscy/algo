package ac

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0

	for l < r {
		t := min(height[l], height[r]) * (r - l)
		ans = max(t, ans)
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return ans

}
