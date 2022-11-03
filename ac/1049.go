package ac

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	//dp := initArray(len(stones)+1, target+1)
	dp := make([]int, target+1)
	for i := 1; i <= len(stones); i++ {
		for j := target; j >= stones[i-1]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i-1]]+stones[i-1])
		}
	}
	return sum - 2*dp[target]
}
