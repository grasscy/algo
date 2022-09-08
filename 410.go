package algo

func splitArray(nums []int, m int) int {
	left, right := nums[0], 0
	for _, v := range nums {
		if left < v {
			left = v
		}
		right += v
	}

	for left < right {
		mid := (left + right) / 2
		d := f(nums, mid)
		if d <= m {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func f(nums []int, x int) int {
	m := 1
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if sum > x {
			m++
			sum = nums[i]
		}
	}
	return m
}
