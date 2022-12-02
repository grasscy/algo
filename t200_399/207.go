package t200_399

var prem map[int][]int
var vis []int

func canFinish(numCourses int, prerequisites [][]int) bool {
	prem = map[int][]int{}
	for _, v := range prerequisites {
		if vv, ok := prem[v[0]]; ok {
			prem[v[0]] = append(vv, v[1])
		} else {
			prem[v[0]] = []int{v[1]}
		}
	}

	for i := 0; i < numCourses; i++ {
		vis = make([]int, numCourses)
		if !d(i) {
			return false
		}
	}

	return true
}

func d(n int) bool {
	if vis[n] == 1 {
		return false
	}
	if vis[n] == -1 {
		return true
	}
	vis[n] = 1
	for _, v := range prem[n] {
		if !d(v) {
			return false
		}
	}
	vis[n] = -1
	return true
}
