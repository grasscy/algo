package t1_199

func candy(ratings []int) int {

	cs := make([]int, len(ratings))
	cs[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			cs[i] = cs[i-1] + 1
		} else {
			cs[i] = 1
			if cs[i-1] >= 2 {
			} else {
				for j := i; j >= 1 && ratings[j-1] > ratings[j] && cs[j-1] <= cs[j]; j-- {
					cs[j-1]++
				}
			}
		}
	}
	ans := 0
	for _, v := range cs {
		ans += v
	}
	return ans
}
