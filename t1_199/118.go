package t1_199

func generate(numRows int) [][]int {

	ans := [][]int{{1}}
	if numRows == 1 {
		return ans
	}
	for i := 1; i < numRows; i++ {
		t := make([]int, i+1)
		t[0] = 1
		t[i] = 1
		for j := 1; j < i; j++ {
			t[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans = append(ans, t)
	}
	return ans
}
