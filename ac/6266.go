package ac

import "math"

var zhi []bool

func smallestValue(n int) int {
	zhi = make([]bool, n+1)
	zhi[1] = true
	zhi[2] = true
	for i := range zhi {
		f := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				f = false
				break
			}
		}
		if f {
			zhi[i] = true
		}
	}

	ans := n
	for !zhi[n] {
		i := s(n)
		if ans > i {
			ans = i
		}
		if n == i {
			break
		}
		n = i
	}
	return ans
}

func s(n int) int {
	if zhi[n] {
		return n
	}
	for i := 2; i < n; i++ {
		if zhi[i] && (n%i == 0) {
			return i + s(n/i)
		}
	}
	return 0
}
