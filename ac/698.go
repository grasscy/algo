package ac

import "sort"

var used []bool

func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) < k {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	used = make([]bool, len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return back(nums, target, 0, 0, k, 0)
}

func back(nums []int, target int, cursum int, n int, k int, idx int) bool {
	if n == k {
		return true
	}
	if cursum == target {
		return back(nums, target, 0, n+1, k, 0)
	}
	for i := idx; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if cursum+nums[i] > target {
			continue
		}
		used[i] = true
		cursum += nums[i]
		if back(nums, target, cursum, n, k, i+1) {
			return true
		}

		used[i] = false
		cursum -= nums[i]

	}
	return false
}
