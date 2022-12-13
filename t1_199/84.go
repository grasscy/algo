package t1_199

func largestRectangleArea(heights []int) int {
	ans := 0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stk := []int{}
	for i, v := range heights {
		if len(stk) == 0 {
			stk = append(stk, i)
			continue
		}
		top := stk[len(stk)-1]
		if v >= heights[top] {
			stk = append(stk, i)
		} else {
			for heights[top] > v { //0 2 1 2 0
				h := heights[top]
				stk = stk[:len(stk)-1]
				l := i - stk[len(stk)-1] - 1
				s := h * l
				if ans < s {
					ans = s
				}
				top = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}

	}

	return ans
}
