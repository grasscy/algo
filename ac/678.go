package ac

func checkValidString(s string) bool {
	stkleft := []int{}
	stkxing := []int{}
	for i, v := range s {
		switch v {
		case '(':
			stkleft = append(stkleft, i)
		case '*':
			stkxing = append(stkxing, i)
		case ')':
			if len(stkleft) > 0 {
				stkleft = stkleft[:len(stkleft)-1]
			} else if len(stkxing) > 0 {
				stkxing = stkxing[:len(stkxing)-1]
			} else {
				return false
			}
		}
	}
	for len(stkleft) > 0 && len(stkxing) > 0 {
		l := stkleft[len(stkleft)-1]
		x := stkxing[len(stkxing)-1]
		stkleft = stkleft[:len(stkleft)-1]
		stkxing = stkxing[:len(stkxing)-1]
		if l > x {
			return false
		}
	}
	return len(stkleft) == 0
}
