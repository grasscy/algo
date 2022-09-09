package ac

func minAddToMakeValid(s string) int {
	need := 0
	ans := 0
	for _, v := range s {
		if string(v) == "(" {
			need++
		}
		if string(v) == ")" {
			need--
			if need == -1 {
				ans++
				need = 0
			}
		}
	}
	return ans + need
}
