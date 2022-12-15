package t1_199

var tos map[string][]string
var ans [][]string
var vis map[string]bool

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	tos = map[string][]string{}
	ans = [][]string{}
	found := false
	vis = map[string]bool{}

	for _, v := range wordList {
		if v == endWord {
			found = true
		}
	}
	if !found {
		return ans
	}

	for i := 0; i < len(wordList); i++ {
		from := wordList[i]
		for j := i + 1; j < len(wordList); j++ {
			to := wordList[j]
			if canTo(from, to) {
				if _, ok := tos[from]; !ok {
					tos[from] = []string{}
				}
				if _, ok := tos[to]; !ok {
					tos[to] = []string{}
				}
				tos[from] = append(tos[from], to)
				tos[to] = append(tos[to], from)
			}
		}
	}
	tos[beginWord] = []string{}
	for _, v := range wordList {
		if canTo(beginWord, v) {
			tos[beginWord] = append(tos[beginWord], v)
		}
	}

	steps := map[string]int{}
	steps[beginWord] = 0
	from := map[string][]string{}
	found = bfs(beginWord, endWord, from, steps)
	if found {
		dfs(beginWord, from, []string{endWord}, endWord)
	}
	return ans
}
func dfs(beginWord string, from map[string][]string, tmp []string, cur string) {
	if cur == beginWord {
		ans = append(ans, append([]string{}, tmp...))
		return
	}
	for _, v := range from[cur] {
		tmp = append([]string{v}, tmp...)
		dfs(beginWord, from, tmp, v)
		tmp = tmp[1:]
	}
}

func bfs(beginWord string, endWord string, from map[string][]string, steps map[string]int) bool {
	found := false
	q := []string{beginWord}
	step := 0
	for len(q) != 0 {
		step++
		size := len(q)

		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			for _, v := range tos[cur] {
				if vv, ok := steps[v]; ok {
					if vv == step {
						from[v] = append(from[v], cur)
					}
				}
				if vis[v] {
					continue
				}
				vis[v] = true
				q = append(q, v)

				if _, ok := from[v]; !ok {
					from[v] = []string{}
				}
				from[v] = append(from[v], cur)
				steps[v] = step
				if v == endWord {
					found = true
				}
			}
		}
		if found {
			break
		}
	}
	return found
}

func canTo(from, to string) bool {
	n := 0
	for i := range from {
		if from[i] != to[i] {
			n++
		}
	}
	return n <= 1
}
