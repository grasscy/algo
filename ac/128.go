package ac

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	ans := 0
	for _, v := range nums {
		if m[v-1] {
			continue
		}
		cur := 0
		for m[v] {
			cur++
			v++
		}
		if ans < cur {
			ans = cur
		}

	}
	return ans
}
