package t1_199

// opt: 位运算
func singleNumber(nums []int) int {
	single := 0
	for _, v := range nums {
		single ^= v
	}
	return single
}
