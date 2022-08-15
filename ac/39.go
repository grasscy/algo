package ac

var r [][]int

func combinationSum(candidates []int, target int) [][]int {
    r = [][]int{}
    f(target, candidates, 0, 0, []int{})
    return r
}

func f(target int, candidates []int, index int, sum int, cur []int) {
    if sum == target {
        r = append(r, append([]int{}, cur...))
        return
    }
    if sum > target {
        return
    }
    for i := index; i < len(candidates); i++ {
        cur = append(cur, candidates[i])
        f(target, candidates, i, sum+candidates[i], cur)
        cur = cur[:len(cur)-1]
    }
}
