package ac

import "math/bits"

func hammingDistance(x int, y int) int {
	r := x ^ y
	return bits.OnesCount(uint(r))
}
