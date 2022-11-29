package pass

func lengthOfLastWord(s string) int {
	pre := 0
	t := 0
	for _, v := range s {
		if v == ' ' {
			if t != 0 {
				pre = t
			}
			t = 0
		} else {
			t++
		}
	}
	if t == 0 {
		return pre
	}
	return t
}
