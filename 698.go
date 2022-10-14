package algo

import "sort"

var ans bool
var vis []bool

func canPartitionKSubsets(nums []int, k int) bool {
	ans = false
	vis = make([]bool, len(nums))
	sort.Ints(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k

}

func b(nums []int, target int, index int, sum int) {
	if index == len(nums) && target == sum {
		ans = true
		return
	}

}
