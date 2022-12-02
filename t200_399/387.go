package t200_399

func firstUniqChar(s string) int {
	m := map[int32]int{}
	for _, v := range s {
		m[v]++
	}
	ans := -1
	for i, v := range s {
		if m[v] == 1 {
			ans = i
			break
		}
	}
	return ans
}
