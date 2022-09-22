package ac

import "sort"

var ans [][]int

func subsetsWithDup(nums []int) [][]int {
	ans = [][]int{}
	sort.Ints(nums)
	b(nums, 0, []int{})
	return ans
}

func b(nums []int, start int, cur []int) {
	ans = append(ans, append([]int{}, cur...))

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		cur = append(cur, nums[i])
		b(nums, i+1, cur)
		cur = cur[:len(cur)-1]
	}
}
