package t200_399

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	ans := []string{}
	start := nums[0]
	end := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			if start == end {
				ans = append(ans, strconv.Itoa(start))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", start, end))
			}
			start = nums[i]
			end = nums[i]
		} else {
			end = nums[i]
		}
	}
	if start == end {
		ans = append(ans, strconv.Itoa(start))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", start, end))
	}
	return ans
}
