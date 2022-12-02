package t1_199

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{}
	}
	if rowIndex == 1 {
		return []int{1}
	}
	var ans []int
	for i := 1; i < rowIndex; i++ {
		t := make([]int, i+1)
		t[0] = 1
		t[i] = 1
		for j := 1; j < i; j++ {
			t[j] = ans[j-1] + ans[j]
		}
		ans = t
	}
	return ans
}
