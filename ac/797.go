package ac

var ans [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	ans = [][]int{}
	dfs(graph, 0, []int{})
	return ans
}

func dfs(graph [][]int, i int, path []int) {
	path = append(path, i)
	if i == len(graph)-1 {
		ans = append(ans, append([]int{}, path...))
	}
	for _, v := range graph[i] {
		dfs(graph, v, path)
	}
	path = path[:len(path)-1]
}
