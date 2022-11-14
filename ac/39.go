package ac

import "sort"

var ans [][]int

func combinationSum(candidates []int, target int) [][]int {
    ans = [][]int{}
    sort.Ints(candidates)
    b(candidates, target, 0, []int{}, 0)
    return ans
}

func b(candidates []int, target int, sum int, tmp []int, index int) {
    if target == sum {
        ans = append(ans, append([]int{}, tmp...))
        return
    }
    if sum > target {
        return
    }
    for i := index; i < len(candidates); i++ {
        if sum+candidates[i] > target {
            return
        }
        b(candidates, target, sum+candidates[i], append(tmp, candidates[i]), i)
    }
}
