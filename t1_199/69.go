package t1_199

func mySqrt(x int) int {
	n := 1
	for n*n <= x {
		n += 1
	}
	return n
}
