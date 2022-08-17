package ac

import "sort"

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] < intervals[j][0] {
            return true
        }
        return false
    })

    r := [][]int{}
    start, end := intervals[0][0], intervals[0][1]
    for i := 0; i < len(intervals); i++ {
        if intervals[i][0] > end {
            r = append(r, []int{start, end})
            start = intervals[i][0]
            end = intervals[i][1]
        } else {
            end = max(end, intervals[i][1])
        }
    }
    r = append(r, []int{start, end})
    return r
}
