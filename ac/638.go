package ac

import (
	"math"
)

var ans int

func shoppingOffers(price []int, special [][]int, needs []int) int {
	ans = math.MaxInt32
	d(price, special, needs, 0, make([]int, len(needs)), 0)
	return ans
}

func d(price []int, special [][]int, needs []int, sum int, buys []int, index int) {
	for i := 0; i < len(needs); i++ {
		if needs[i] < buys[i] {
			return
		}
	}

	tmpsum := sum
	for i := 0; i < len(needs); i++ {
		n := needs[i] - buys[i]
		tmpsum += price[i] * n
	}
	if tmpsum < ans {
		ans = tmpsum
	}

	for i := index; i < len(special); i++ {
		for j := 0; j < len(buys); j++ {
			buys[j] += special[i][j]
		}
		sum += special[i][len(buys)]

		d(price, special, needs, sum, buys, i)

		sum -= special[i][len(buys)]
		for j := 0; j < len(buys); j++ {
			buys[j] -= special[i][j]
		}
	}

}
