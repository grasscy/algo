package algo

func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    // opt: 滚动数组
    //first := nums[0]
    //second := max(nums[0], nums[1])
    //for j := 2; j < len(nums; j++ {
    //  first, second = second, max(first + nums[i], second)
    //
    //}

    dp := make([]int, len(nums))
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])
    for i := 1; i < len(nums); i++ {
        dp[i] = max(dp[i-2]+nums[i], dp[i-1])
    }

    return dp[len(nums)-1]
}
