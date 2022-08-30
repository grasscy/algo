package ac

func subarraySum(nums []int, k int) int {

	m := map[int]int{}
	m[0] = 1

	presum := 0
	ans := 0

	for i := 0; i < len(nums); i++ {
		presum += nums[i]
		ans += m[presum-k]
		m[presum]++
	}

	return ans
}
