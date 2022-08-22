package ac

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
    in := make([]int, numCourses)
    m := map[int][]int{}

    for i := 0; i < len(prerequisites); i++ {
        in[prerequisites[i][0]]++
        if _, ok := m[prerequisites[i][1]]; ok {
            m[prerequisites[i][1]] = append(m[prerequisites[i][1]], prerequisites[i][0])
        } else {
            m[prerequisites[i][1]] = []int{prerequisites[i][0]}
        }
    }
    queue := []int{}
    for i, v := range in {
        if v == 0 {
            queue = append(queue, i)
        }
    }

    count := 0

    for len(queue) > 0 {
        c := queue[0]
        fmt.Println("c: ", c)
        queue = queue[1:]
        count++
        if v, ok := m[c]; ok {
            for _, cc := range v {
                in[cc]--
                if in[cc] == 0 {
                    queue = append(queue, cc)
                }
            }
        }
        fmt.Println()
        fmt.Println(queue)
    }

    return count == numCourses

}
