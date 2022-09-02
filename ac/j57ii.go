package ac

func findContinuousSequence(target int) [][]int {
	ans := [][]int{}
	i, j := 1, 1
	sum := 0
	for i <= target/2 {
		if sum == target {
			a := []int{}
			for k := i; k < j; k++ {
				a = append(a, k)
			}
			ans = append(ans, a)
			sum -= i
			i++
		} else if sum < target {
			sum += j
			j++
		} else {
			sum -= i
			i++
		}
	}

	return ans
}
