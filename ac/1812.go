package ac

func squareIsWhite(coordinates string) bool {
	// 01 10 hei,00 11bai
	j := coordinates[0] - 'a'
	i := coordinates[1] - '0'
	if i%2 != j%2 {
		return false
	}
	return true
}
