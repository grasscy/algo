package ac

func beautySum(s string) int {
	ans := 0
	for i := range s {
		cnt := map[string]int{}

		for j := i; j < len(s); j++ {
			cnt[string(s[j])]++
			mmax := -1
			mmin := 555
			for _, v := range cnt {
				if v > mmax {
					mmax = v
				}
				if v < mmin && v != 0 {
					mmin = v
				}
			}
			ans += mmax - mmin
		}
	}

	return ans
}
