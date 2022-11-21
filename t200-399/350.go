package t200_399

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, v := range nums1 {
		m[v]++
	}
	ans := []int{}
	for _, v := range nums2 {
		if m[v] != 0 {
			ans = append(ans, v)
			m[v]--
		}
	}
	return ans
}
