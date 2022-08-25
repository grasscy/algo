package ac

func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	right := len(arr) - 1

	r := []int{}
	for i := 0; i < len(arr)-k; i++ {
		if abs(arr[left]-x) > abs(arr[right]-x) {
			left++
		} else {
			right--
		}
	}

	for i := left; i <= right; i++ {
		r = append(r, arr[i])
	}
	return r
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
