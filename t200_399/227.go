package t200_399

import "strconv"

//func calculate(s string) (ans int) {
//    stack := []int{}
//    preSign := '+'
//    num := 0
//    for i, ch := range s {
//        isDigit := '0' <= ch && ch <= '9'
//        if isDigit {
//            num = num*10 + int(ch-'0')
//        }
//        if !isDigit && ch != ' ' || i == len(s)-1 {
//            switch preSign {
//            case '+':
//                stack = append(stack, num)
//            case '-':
//                stack = append(stack, -num)
//            case '*':
//                stack[len(stack)-1] *= num
//            default:
//                stack[len(stack)-1] /= num
//            }
//            preSign = ch
//            num = 0
//        }
//    }
//    for _, v := range stack {
//        ans += v
//    }
//    return
//}
func calculate(s string) int {
	stkf := []string{}
	stkn := []int{}

	nn := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			nn += string(s[i])
		} else if s[i] == '+' || s[i] == '-' {
			atoi, _ := strconv.Atoi(nn)
			stkn = append(stkn, atoi)
			nn = ""
			stkf = append(stkf, string(s[i]))
		} else {
			atoi, _ := strconv.Atoi(nn)

			bb := ""
			k := i
			i++
			for ; i < len(s); i++ {
				if s[i] == ' ' {
					continue
				}
				if s[i] >= '0' && s[i] <= '9' {
					bb += string(s[i])
				} else {
					break
				}
			}
			i--
			b, _ := strconv.Atoi(bb)

			var aa int
			if s[k] == '*' {
				aa = atoi * b
			} else {
				aa = atoi / b
			}
			nn = strconv.Itoa(aa)
		}
	}
	atoi, _ := strconv.Atoi(nn)
	stkn = append(stkn, atoi)

	for len(stkf) != 0 {
		f := stkf[0]
		stkf = stkf[1:]
		b := stkn[1]
		a := stkn[0]

		if f == "+" {
			stkn[1] = a + b
		} else {
			stkn[1] = a - b
		}
		stkn = stkn[1:]

	}

	return stkn[0]

}
