package ac

func countQuadruplets(nums []int) int {
	ans := 0
	for i := 3; i < len(nums); i++ {
		for j := 0; j < i-2; j++ {
			for k := j + 1; k < i-1; k++ {
				for m := k + 1; m < i; m++ {
					if nums[j]+nums[k]+nums[m] == nums[i] {
						ans++
					}
				}
			}
		}
	}
	return ans
}
