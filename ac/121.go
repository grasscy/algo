package ac

import "math"

func maxProfit(prices []int) int {
    max := 0
    minp := math.MaxInt32
    for _, v := range prices {
        if v-minp > max {
            max = v - minp
        }

        if v < minp {
            minp = v
        }
    }
    return max
}
