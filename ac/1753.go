package ac

import "sort"

func maximumScore(a int, b int, c int) int {
	ans := 0
	for {
		ns := append([]int{}, a, b, c)
		sort.Ints(ns)
		a, b, c = ns[0], ns[1], ns[2]
		if a == 0 && b == 0 {
			break
		}
		if a+b > c {
			a--
			b--
			ans++
		} else {
			ans += a + b
			break
		}
	}
	return ans

}
