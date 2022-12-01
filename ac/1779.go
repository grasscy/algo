package ac

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	mmin := math.MaxFloat64
	ans := -1
	for i, v := range points {
		if v[0] == x || v[1] == y {
			h := math.Abs(float64(x-v[0])) + math.Abs(float64(y-v[1]))
			if h < mmin {
				ans = i
				mmin = h
			}
		}
	}
	return ans
}
