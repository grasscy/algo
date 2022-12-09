package ac

import "math"

var ns []float64
var vis []bool

//func checkPowersOfThree(n int) bool {
//    for ; n > 0; n /= 3 {
//        if n%3 == 2 {
//            return false
//        }
//    }
//    return true
//}

func checkPowersOfThree(n int) bool {
	// 1 3 9 7 1
	ns = []float64{}
	for i := 0; math.Pow(3, float64(i)) <= float64(n); i++ {
		ns = append(ns, math.Pow(3, float64(i)))
	}
	vis = make([]bool, len(ns))
	return b(n, 0, 0)
}

func b(n int, sum int, index int) bool {
	if sum == n {
		return true
	}
	if sum > n {
		return false
	}
	for i := index; i < len(ns); i++ {
		if b(n, sum+int(ns[i]), i+1) {
			return true
		}
	}
	return false
}
