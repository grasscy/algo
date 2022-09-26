package ac

func plusOne(digits []int) []int {
	jin := 0
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i]

		if i == len(digits)-1 {
			n += 1
		}
		if n+jin > 9 {
			digits[i] = (n + jin) % 10
			jin = (n + jin) / 10
		} else {
			digits[i] = n + jin
			jin = 0
		}
	}
	ans := []int{}
	if jin != 0 {
		ans = append([]int{1}, digits...)
	} else {
		ans = append([]int{}, digits...)
	}
	return ans

}
