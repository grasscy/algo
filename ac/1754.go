package ac

func largestMerge(word1 string, word2 string) string {
	m := ""
	for len(word1) != 0 && len(word2) != 0 {
		if word1[0] > word2[0] {
			m += string(word1[0])
			word1 = word1[1:]
		} else if word1[0] < word2[0] {
			m += string(word2[0])
			word2 = word2[1:]
		} else {

			f := true
			i := 1
			for ; i < len(word1) && i < len(word2); i++ {
				if word1[i] == word2[i] {
					continue
				}
				if word1[i] > word2[i] {
					break
				} else {
					f = false
					break
				}
			}
			if i == len(word1) {
				m += string(word2[0])
				word2 = word2[1:]
			} else if i == len(word2) {
				m += string(word1[0])
				word1 = word1[1:]
			} else if f {
				m += string(word1[0])
				word1 = word1[1:]
			} else {
				m += string(word2[0])
				word2 = word2[1:]
			}
		}
	}
	for len(word1) != 0 {
		m += string(word1[0])
		word1 = word1[1:]
	}
	for len(word2) != 0 {
		m += string(word2[0])
		word2 = word2[1:]
	}
	return m
}
