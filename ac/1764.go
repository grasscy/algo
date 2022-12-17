package ac

func canChoose(groups [][]int, nums []int) bool {
	start := 0
	n := 0
	for _, v := range groups {
		vs := 0
		k := 0
		i := start + k
		for k < len(nums) && i < len(nums) {
			if nums[i] == v[vs] {
				vs++
				i++
			} else {
				k++
				i = start + k
				vs = 0
			}
			if vs == len(v) {
				break
			}
		}

		if vs != len(v) {
			return false
		}
		start = i
		n++
	}
	return n == len(groups)
}
