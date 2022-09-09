package ac

func minInsertions(s string) int {
	need := 0
	ans := 0
	for _, v := range s {
		if string(v) == "(" {
			need += 2
			if need%2 == 1 {
				ans++
				need--
			}
		}
		if string(v) == ")" {
			need -= 1
			if need == -1 {
				ans++
				need = 1
			}
		}
	}

	return ans + need
}
