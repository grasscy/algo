package ac

func longestSubsequence(arr []int, difference int) int {
	ans := 1
	dp := map[int]int{}
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return ans

	//dp := make([]int, len(arr))
	//dp[0] = 1
	//for i := 1; i < len(dp); i++ {
	//    dp[i] = 1
	//    for j := 0; j < i; j++ {
	//        if arr[j]+difference == arr[i] {
	//            dp[i] = max(dp[i], dp[j]+1)
	//        }
	//    }
	//}
	//mm := 1
	//for i := 0; i < len(dp); i++ {
	//    mm = max(mm, dp[i])
	//}
	//return mm
}
