package algo

func busyStudent(startTime []int, endTime []int, queryTime int) int {
    r := 0
    for i := 0; i < len(startTime); i++ {
        if startTime[i] <= queryTime && endTime[i] >= queryTime {
            r++
        }
    }
    return r
}
