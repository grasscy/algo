package ac

func minimumSize(nums []int, maxOperations int) int {
	mmax := -1
	for _, v := range nums {
		if v > mmax {
			mmax = v
		}
	}
	l, r := 1, mmax
	for l < r {
		mid := (l + r) / 2
		if check(mid, nums, maxOperations) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func check(mid int, nums []int, ops int) bool {
	for _, v := range nums {
		ops -= (v - 1) / mid
	}
	return ops >= 0
}
