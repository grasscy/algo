package algo

func maxSubArray(nums []int) int {
    max := nums[0]
    pre := nums[0]
    for i := 1; i < len(nums); i++ {
        if pre > 0 {
            pre = pre + nums[i]
        } else {
            pre = nums[i]
        }
        if pre > max {
            max = pre
        }
    }

    return max
}
