package algo

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    m := map[int]bool{}
    for _, v := range nums {
        m[v] = true
    }
    r := 1
    for _, v := range nums {
        if!m[v-1]{
            for i := 1; ; i++ {
                _, ok := m[v+i]
                if !ok {
                    r = max(r, i)
                    break
                }
            }
        }
    }
    return r
}
