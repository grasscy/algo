package ac

func checkIfPangram(sentence string) bool {
	n := make([]int, 26)
	for _, v := range sentence {
		n[v-'a']++
	}
	for _, v := range n {
		if v == 0 {
			return false
		}
	}
	return true
}
