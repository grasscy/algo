package pass

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	f := make([]int, m)
	if obstacleGrid[0][0] == 1 {
		f[0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j-1 > 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
	//
	//m := len(obstacleGrid)
	//n := len(obstacleGrid[0])
	//
	//dp := make([][]int, m)
	//for i := 0; i < m; i++ {
	//    dp[i] = make([]int, n)
	//}
	//i := 0
	//for ; i < m; i++ {
	//    if obstacleGrid[i][0] == 0 {
	//        dp[i][0] = 1
	//    } else {
	//        break
	//    }
	//}
	//for ; i < m; i++ {
	//    dp[i][0] = 0
	//}
	//i = 0
	//for ; i < n; i++ {
	//    if obstacleGrid[0][i] == 0 {
	//        dp[0][i] = 1
	//    } else {
	//        break
	//    }
	//}
	//for ; i < n; i++ {
	//    dp[0][i] = 0
	//}
	//
	//for i := 1; i < m; i++ {
	//    for j := 1; j < n; j++ {
	//        if obstacleGrid[i-1][j] == 1 {
	//            dp[i-1][j] = 0
	//        }
	//        if obstacleGrid[i][j-1] == 1 {
	//            dp[i][j-1] = 0
	//        }
	//        if obstacleGrid[i][j] == 1 {
	//            dp[i][j] = 0
	//        } else {
	//            dp[i][j] = dp[i-1][j] + dp[i][j-1]
	//        }
	//    }
	//}
	//return dp[m-1][n-1]
}
