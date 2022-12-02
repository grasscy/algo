package ac

import "math"

func minOperations(boxes string) []int {
	had := []int{}
	for i, v := range boxes {
		if v == '1' {
			had = append(had, i)
		}
	}
	ans := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		sum := 0
		for _, v := range had {
			sum += int(math.Abs(float64(i - v)))
		}
		ans[i] = sum
	}
	return ans
}
