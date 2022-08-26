package ac

func canPartition(nums []int) bool {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
	}

	dp[0][0] = true
	if nums[0] <= target {
		dp[0][nums[0]] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < target+1; j++ {
			if nums[i] == j {
				dp[i][j] = true
			} else if nums[i] < j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}

			if dp[i][target] {
				return true
			}
		}
	}

	//for (int i = 1; i < len; i++) {
	//    for (int j = target; nums[i] <= j; j--) {
	//        if (dp[target]) {
	//            return true;
	//        }
	//        dp[j] = dp[j] || dp[j - nums[i]];
	//    }
	//}
	//return dp[target];

	return dp[len(nums)-1][target]
}
