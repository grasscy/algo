package t1_199

import "strings"

func isNumber(s string) bool {
	if strings.Count(s, ".") > 1 || strings.Count(s, "e") > 1 || strings.Count(s, "E") > 1 {
		return false
	}

	for i, v := range s {
		if v == '-' || v == '+' {
			if i != 0 || i == len(s)-1 {
				return false
			}

		} else if v == '.' {
			if i == len(s)-1 {
				if i > 0 && s[i-1] >= '0' && s[i-1] <= '9' {
					continue
				}
				return false
			}
			if i > 0 {
				if s[i-1] >= '0' && s[i-1] <= '9' {
					continue
				}
				if s[i-1] == '-' || s[i-1] == '+' {
					continue
				}
			}
			if i < len(s)-1 {
				if s[i+1] >= '0' && s[i+1] <= '9' {
					continue
				}
			}
			return false
		} else if v == 'e' || v == 'E' {
			if i == 0 || s[i-1] == '-' || s[i-1] == '+' {
				return false
			}
			if i == len(s)-1 {
				return false
			}
			if isZh(s[i+1:]) {
				return true
			} else {
				return false
			}
		} else if v >= '0' && v <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
func isZh(s string) bool {
	for i, v := range s {
		if v == '+' || v == '-' {
			if i != 0 {
				return false
			}
			if i == len(s)-1 {
				return false
			}
		} else if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
