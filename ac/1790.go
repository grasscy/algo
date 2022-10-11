package ac

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	diffs := []int{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffs = append(diffs, i)
			if len(diffs) > 2 {
				return false
			}
		}
	}
	if len(diffs) != 2 {
		return false
	}
	if s1[diffs[0]] == s2[diffs[1]] && s1[diffs[1]] == s2[diffs[0]] {
		return true
	}
	return false

}
