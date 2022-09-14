package ac

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr) / 20
	sum := 0
	for i := n; i < len(arr)-n; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr)-2*n)
}
