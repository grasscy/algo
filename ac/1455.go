package ac

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	split := strings.Split(sentence, " ")
	for i := 0; i < len(split); i++ {
		if strings.HasPrefix(split[i], searchWord) {
			return i + 1
		}
	}
	return -1
}
