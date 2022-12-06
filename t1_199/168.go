package t1_199

func convertToTitle(columnNumber int) string {
	ans := ""
	for columnNumber != 0 {
		if columnNumber%26 == 0 {
			ans = "Z" + ans
			columnNumber = columnNumber/26 - 1
		} else {
			ans = string(rune(columnNumber%26+'A'-1)) + ans
			columnNumber = columnNumber / 26
		}
	}
	return ans
}
