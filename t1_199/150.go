package t1_199

import "strconv"

func evalRPN(tokens []string) int {
	nstk := []int{}
	for _, v := range tokens {
		switch v {
		case "+":
			a := nstk[len(nstk)-1]
			nstk = nstk[:len(nstk)-1]
			b := nstk[len(nstk)-1]
			nstk = nstk[:len(nstk)-1]
			nstk = append(nstk, a+b)
		case "-":
			a := nstk[len(nstk)-1]
			nstk = nstk[:len(nstk)-1]
			b := nstk[len(nstk)-1]
			nstk = nstk[:len(nstk)-1]
			nstk = append(nstk, b-a)
		case "*":
			a := nstk[len(nstk)-1]
			nstk = nstk[:len(nstk)-1]
			b := nstk[len(nstk)-1]
			nstk = nstk[:len(nstk)-1]
			nstk = append(nstk, a*b)
		case "/":
			a := nstk[len(nstk)-1]
			nstk = nstk[:len(nstk)-1]
			b := nstk[len(nstk)-1]
			nstk = nstk[:len(nstk)-1]
			nstk = append(nstk, b/a)
		default:
			atoi, _ := strconv.Atoi(v)
			nstk = append(nstk, atoi)
		}
	}
	return nstk[0]
}
