package algo

import "math"

func minMoves(nums []int, k int) int {
	p := []int{}
	for i, v := range nums {
		if v != 0 {
			p = append(p, i-len(p))
		}
	}
	m := len(p)
	s := make([]int, m+1) // p 的前缀和
	for i, v := range p {
		s[i+1] = s[i] + v
	}
	ans := math.MaxInt32
	for i, v := range s[:m-k+1] { // p[i] 到 p[i+k-1] 中所有数到 p[i+k/2] 的距离之和，取最小值
		ans = min(ans, v+s[i+k]-s[i+k/2]*2-p[i+k/2]*(k%2))
	}
	return ans
}

//
//func minMoves(nums []int, k int) int {
//    zeros := []int{}
//    i := 0
//    for ; i < len(nums) && nums[i] == 0; i++ {
//    }
//    for i < len(nums) {
//        j := i + 1
//        for j < len(nums) && nums[j] == 0 {
//            j++
//        }
//        if j < len(nums) {
//            zeros = append(zeros, j-i-1)
//        }
//        i = j
//    }
//
//    cost := 0
//    l, r := 0, k-2
//    for i := l; i <= r; i++ {
//        cost += zeros[i] * min(i+1, r-i+1)
//    }
//
//    ans := cost
//    pre := []int{0}
//    for _, v := range zeros {
//        pre = append(pre, pre[len(pre)-1]+v)
//    }
//
//    i = 1
//    j := i + k - 2
//    for ; j < len(zeros); {
//        mid := (i + j) / 2
//        cost -= GetRangeSum(i-1, mid-1, pre)
//        cost += GetRangeSum(mid+k%2, j, pre)
//        ans = min(ans, cost)
//        i++
//        j++
//    }
//    return ans
//}
//
//func GetRangeSum(left, right int, pre []int) int {
//    return pre[right+1] - pre[left]
//}
