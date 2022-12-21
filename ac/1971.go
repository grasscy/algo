package ac

var to map[int][]int
var vis []bool

func validPath(n int, edges [][]int, source int, destination int) bool {
	vis = make([]bool, n)
	to = map[int][]int{}
	for _, v := range edges {
		to[v[0]] = append(to[v[0]], v[1])
		to[v[1]] = append(to[v[1]], v[0])
	}
	return d(source, destination)
}

func d(source int, destination int) bool {
	if source == destination {
		return true
	}
	for _, v := range to[source] {
		if vis[v] {
			continue
		}
		vis[v] = true
		if d(v, destination) {
			return true
		}
	}
	return false
}
