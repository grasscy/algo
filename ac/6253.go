package ac

import "strings"

func isCircularSentence(sentence string) bool {
	split := strings.Split(sentence, " ")
	if len(split) == 1 {
		return split[0][0] == split[0][len(split[0])-1]
	}
	pre := split[0][len(split[0])-1]
	for i := 1; i < len(split); i++ {
		if split[i][0] != pre {
			return false
		}
		pre = split[i][len(split[i])-1]
	}
	return pre == split[0][0]

}
