package t1_199

import "sort"

var ans [][]int

func subsetsWithDup(nums []int) [][]int {
	ans = [][]int{}
	sort.Ints(nums)
	b(nums, []int{}, 0)
	return ans
}

func b(nums []int, tmp []int, index int) {
	ans = append(ans, append([]int{}, tmp...))
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		tmp = append(tmp, nums[i])
		b(nums, tmp, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}
