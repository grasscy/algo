package ac

import "math/rand"

func findKthLargest(nums []int, k int) int {
	return qSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func qSelect(a []int, l, r int, index int) int {
	q := partition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return qSelect(a, q+1, r, index)
	}
	return qSelect(a, l, q-1, index)
}

func partition(a []int, l, r int) int {
	q := rand.Intn(r-l+1) + l
	a[r], a[q] = a[q], a[r]

	i := l - 1
	for j := l; j < r; j++ {
		if a[j] < a[r] {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}
