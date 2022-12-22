package t1_199

import "sort"

func maxPoints(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		a, b := points[i], points[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	if len(points) == 1 {
		return 1
	}
	ans := 2
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a, b := points[i], points[j]
			k := float64(b[1]-a[1]) / float64(b[0]-a[0])
			s := 2
			for jj := j + 1; jj < len(points); jj++ {
				c := points[jj]
				kk := float64(c[1]-a[1]) / float64(c[0]-a[0])
				if kk == k {
					s++
				}
			}
			ans = max(ans, s)
		}
	}
	return ans

}
