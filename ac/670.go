package ac

import "sort"

func maximumSwap(num int) int {
	nums := []int{}
	for num/10 != 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	nums = append(nums, num)
	raw := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		raw = append(raw, nums[i])
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i, v := range raw {
		if v != nums[i] {
			maxi := 0
			for j := i + 1; j < len(raw); j++ {
				if raw[j] == nums[i] {
					maxi = j
				}
			}
			raw[i], raw[maxi] = raw[maxi], raw[i]
			break
		}
	}

	ans := 0
	for _, v := range raw {
		ans = ans*10 + v
	}

	return ans
}
