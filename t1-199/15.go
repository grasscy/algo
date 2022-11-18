package pass

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j+1 < len(nums) && nums[j] == nums[j+1] {
					j++
				}
				for k-1 > 0 && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}

		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}
