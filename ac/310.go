package ac

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	ans := []int{}
	tos := map[int][]int{}
	degree := make([]int, n)
	for _, v := range edges {
		if _, ok := tos[v[0]]; !ok {
			tos[v[0]] = []int{}
		}
		tos[v[0]] = append(tos[v[0]], v[1])

		if _, ok := tos[v[1]]; !ok {
			tos[v[1]] = []int{}
		}
		tos[v[1]] = append(tos[v[1]], v[0])
		degree[v[0]]++
		degree[v[1]]++
	}
	queue := []int{}
	for i, v := range degree {
		if v == 1 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		size := len(queue)
		ans = []int{}
		for i := 0; i < size; i++ {
			t := queue[0]
			queue = queue[1:]
			ans = append(ans, t)
			if to, ok := tos[t]; ok {
				for _, v := range to {
					degree[v]--
					if degree[v] == 1 {
						queue = append(queue, v)
					}
				}
			}
		}
	}
	return ans
}

//var ans map[int]int
//
////var vis []bool
//var tos map[int][]int
//
//func findMinHeightTrees(n int, edges [][]int) []int {
//	ans := map[int]int{}
//	//vis = make([]bool, n)
//	tos = map[int][]int{}
//	for _, v := range edges {
//		if _, ok := tos[v[0]]; !ok {
//			tos[v[0]] = []int{}
//		}
//		tos[v[0]] = append(tos[v[0]], v[1])
//
//		if _, ok := tos[v[1]]; !ok {
//			tos[v[1]] = []int{}
//		}
//		tos[v[1]] = append(tos[v[1]], v[0])
//	}
//	mm := n
//	for i := 0; i < n; i++ {
//		vis := make([]bool, n)
//		ans[i] = h(i, vis)
//		if mm > ans[i] {
//			mm = ans[i]
//		}
//	}
//	ret := []int{}
//	for k, v := range ans {
//		if v == mm {
//			ret = append(ret, k)
//		}
//	}
//	return ret
//}
//
//func h(n int, vis []bool) int {
//	vis[n] = true
//	v, ok := tos[n]
//	if !ok {
//		return 0
//	}
//	high := 0
//	for _, vv := range v {
//		if vis[vv] {
//			continue
//		}
//		hh := h(vv, vis) + 1
//		if high < hh {
//			high = hh
//		}
//	}
//	return high
//}
