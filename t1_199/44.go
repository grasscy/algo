package t1_199

func isMatch(s string, p string) bool {

	if len(s) == 0 && len(p) == 0 {
		return true
	}
	//
	//index := 0
	//for ; index < len(p) && p[index] == '*'; index++ {
	//}
	//
	//if index > 0 && index == len(p) {
	//    return true
	//}
	//
	//if len(s) == 0 || len(p) == 0 {
	//    return false
	//}

	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true

	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}

	//for i := 0; i < index; i++ {
	//    dp[0][i] = true
	//}
	//
	//if s[0] == p[index] || p[index] == '?' {
	//    dp[0][index] = true
	//
	//    for i := index + 1; i < len(p) && p[i] == '*'; i++ {
	//        dp[0][i] = true
	//    }
	//}
	//
	//if index > 0 {
	//    for i := 0; i < len(s); i++ {
	//        dp[i][0] = true
	//    }
	//}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			//ss := string(s[i])
			//pp := string(p[j])
			//fmt.Println(ss, pp)
			if s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				//for k := i; k < len(s); k++ {
				//    dp[k][j] = dp[k][j] || dp[i-1][j-1] || dp[i][j-1]
				//}
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}

	return dp[len(s)][len(p)]
}
