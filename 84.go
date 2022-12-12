package algo

func largestRectangleArea(heights []int) int {
	ans := 0
	stk := []int{}
	left := make([]int, len(heights))
	for i, v := range heights {
		if v > ans {
			ans = v
		}
		if len(stk) == 0 {
			stk = append(stk, i)
		} else {
			if v < heights[stk[len(stk)-1]] {
				stk = append(stk, i)
			} else {
				for len(stk) != 0 {
					top := stk[len(stk)-1]
					stk = stk[:len(stk)-1]

					left[top] = heights[top] * len(stk)
				}

				stk = append(stk, i)
			}
		}
	}

	return ans
}
