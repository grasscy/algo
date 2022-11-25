package jianzhioffer

var ans [][]int

func subsets(nums []int) [][]int {
	ans = [][]int{}
	b(nums, 0, []int{})
	return ans
}

func b(nums []int, index int, tmp []int) {
	ans = append(ans, append([]int{}, tmp...))

	for i := index; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		b(nums, i+1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}
