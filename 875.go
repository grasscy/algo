package algo

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, int(10e9+1)
	for left < right {
		mid := (left + right) / 2
		if f(piles, mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func f(piles []int, x int) int {
	hours := 0
	for i := 0; i < len(piles); i++ {
		hours += piles[i] / x
		if piles[i]%x != 0 {
			hours++
		}
	}
	return hours
}
