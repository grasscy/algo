package ac

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	dp := initArray(n, n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dp[i][j] = 2
		}
	}

	mm := 0
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[j]-arr[i] > arr[i] {
				break
			}
			for k := i - 1; k >= 0; k-- {
				if arr[k]+arr[i] == arr[j] {
					dp[i][j] = max(dp[i][j], dp[k][i]+1)
					break
				}
			}
			mm = max(mm, dp[i][j])
		}
	}
	if mm <= 2 {
		return 0
	}
	return mm
}
