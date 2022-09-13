package ac

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	stk := []int{}

	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	for i := 0; i < n*2-1; i++ {
		for len(stk) > 0 && nums[i%n] > nums[stk[len(stk)-1]] {
			t := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			ans[t] = nums[i%n]
		}
		stk = append(stk, i%n)
	}
	return ans
}
