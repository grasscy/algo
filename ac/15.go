package ac

import (
    "sort"
)

func threeSum(nums []int) [][]int {
    r := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        j, k := i+1, len(nums)-1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum == 0 {
                r = append(r, []int{nums[i], nums[j], nums[k]})
                for j < k && nums[j] == nums[j+1] {
                    j++
                }
                for j < k && nums[k] == nums[k-1] {
                    k--
                }
                j++
                k--
            } else if sum < 0 {
                j++
            } else {
                k--
            }
        }

    }

    return r
}
