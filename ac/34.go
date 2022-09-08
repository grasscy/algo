package ac

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	start, end := 0, 0
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		start = -1
	} else {
		start = left
	}

	left, right = 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left-1 < 0 || nums[left-1] != target {
		end = -1
	} else {
		end = left - 1
	}

	return []int{start, end}
}
