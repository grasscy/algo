package t200_399

func isHappy(n int) bool {
	m := map[int]bool{}
	for {
		t := 0
		nn := n
		for nn != 0 {
			t += (nn % 10) * (nn % 10)
			nn /= 10
		}
		if t == 1 {
			return true
		}
		if m[t] {
			return false
		}
		m[t] = true
		n = t
	}
}
