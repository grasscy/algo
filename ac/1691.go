package ac

import (
	"encoding/json"
	"fmt"
	"sort"
)

func maxHeight(cuboids [][]int) int {
	for i := 0; i < len(cuboids); i++ {
		sort.Ints(cuboids[i])
	}
	sort.Slice(cuboids, func(i, j int) bool {
		a := cuboids[i]
		b := cuboids[j]
		return a[2] > b[2] || (a[2] == b[2] && a[1] > b[1]) || (a[2] == b[2] && a[1] == b[1] && a[0] > b[0])
		//return a[0] > b[0] || a[0] == b[0] && (a[1] > b[1] || a[1] == b[1] && a[2] > b[2])
	})
	marshal, _ := json.Marshal(cuboids)
	fmt.Println(string(marshal))
	dp := make([]int, len(cuboids))
	dp[0] = cuboids[0][2]

	for i := 1; i < len(cuboids); i++ {
		a := cuboids[i]
		dp[i] = a[2]
		for j := 0; j < i; j++ {
			if can(a, cuboids[j]) {
				dp[i] = max(dp[i], dp[j]+a[2])
			}
		}
	}
	ans := 0
	for _, v := range dp {
		if ans < v {
			ans = v
		}
	}
	return ans
}

func can(a, b []int) bool {
	return a[0] <= b[0] && a[1] <= b[1] && a[2] <= b[2]
}
