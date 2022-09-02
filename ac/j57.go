package ac

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for {
		if nums[i]+nums[j] == target {
			return []int{nums[i], nums[j]}
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			j--
		}
	}
}
