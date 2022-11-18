package pass

func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i+1 < len(s) {
				if s[i+1] == 'V' || s[i+1] == 'X' {
					res -= 2
				}
			}
			res += 1
		case 'V':
			res += 5
		case 'X':
			if i+1 < len(s) {
				if s[i+1] == 'L' || s[i+1] == 'C' {
					res -= 20
				}
			}
			res += 10
		case 'L':
			res += 50
		case 'C':
			if i+1 < len(s) {
				if s[i+1] == 'D' || s[i+1] == 'M' {
					res -= 200
				}
			}
			res += 100
		case 'D':
			res += 500
		case 'M':
			res += 1000
		}
	}
	return res
}
