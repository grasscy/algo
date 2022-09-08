package algo

func shipWithinDays(weights []int, days int) int {

	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}

	for left < right {
		mid := (left + right) / 2
		d := f(weights, mid)
		if d == days {
			right = mid
		} else if d < days {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func f(weights []int, x int) int {
	days := 1
	sum := weights[0]
	for i := 1; i < len(weights); i++ {
		sum += weights[i]
		if sum > x {
			days++
			sum = weights[i]
		}
	}
	return days
}
