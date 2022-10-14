package ac

var ans [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	ans = [][]int{}
	dfs(graph, []int{0}, 0)
	return ans
}

func dfs(graph [][]int, tmp []int, index int) {
	if len(graph)-1 == index {
		ans = append(ans, append([]int{}, tmp...))
		return
	}

	for i := 0; i < len(graph[index]); i++ {
		tmp = append(tmp, graph[index][i])
		dfs(graph, tmp, graph[index][i])
		tmp = tmp[:len(tmp)-1]
	}
}
