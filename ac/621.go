package ac

func leastInterval(tasks []byte, n int) int {
	cnt := map[byte]int{}
	for _, v := range tasks {
		cnt[v]++
	}

	maxExec, lastCnt := 0, 0
	for _, c := range cnt {
		if c > maxExec {
			maxExec = c
			lastCnt = 1
		} else if c == maxExec {
			lastCnt++
		}

	}
	return max(len(tasks), (maxExec-1)*(n+1)+lastCnt)
}
