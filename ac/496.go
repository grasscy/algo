package ac

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ints := next(nums2)
	m := map[int]int{}
	for i, v := range ints {
		m[nums2[i]] = v
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = m[nums1[i]]
	}
	return res
}

func next(nums []int) []int {
	res := make([]int, len(nums))
	stk := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stk) != 0 && stk[len(stk)-1] <= nums[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) != 0 {
			res[i] = stk[len(stk)-1]
		} else {
			res[i] = -1
		}
		stk = append(stk, nums[i])
	}
	return res
}
