package ac

func similarPairs(words []string) int {
	ans := 0
	for i := 0; i < len(words); i++ {
		mi := map[int32]bool{}
		for _, v := range words[i] {
			mi[v] = true
		}

		for j := i + 1; j < len(words); j++ {
			mj := map[int32]bool{}
			for _, v := range words[j] {
				mj[v] = true
			}
			if len(mi) == len(mj) {
				f := true
				for k, _ := range mi {
					if !mj[k] {
						f = false
						break
					}
				}
				if f {
					ans++
				}
			}
		}
	}
	return ans
}
