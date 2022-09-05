package ac

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	for _, v := range bookings {
		diff[v[0]] += v[2]
		if v[1]+1 < n+1 {
			diff[v[1]+1] -= v[2]
		}
	}
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i-1] + diff[i]
	}
	ans = ans[1:]
	return ans
}
