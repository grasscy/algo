package ac

func secondHighest(s string) int {
	mmax := -1
	mmax2 := -1
	for _, v := range s {
		if '0' <= v && '9' >= v {
			i := int(v - '0')
			if i > mmax {
				mmax2 = mmax
				mmax = i
			} else if i < mmax && i > mmax2 {
				mmax2 = i
			}
		}
	}
	return mmax2
}
