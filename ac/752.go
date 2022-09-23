package ac

func openLock(deadends []string, target string) int {
	ans := 0
	vis := map[string]bool{}
	deads := map[string]bool{}
	for _, v := range deadends {
		deads[v] = true
	}
	q := []string{}
	q = append(q, "0000")
	for len(q) != 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[0]
			q = q[1:]
			if deads[cur] {
				continue
			}
			if cur == target {
				return ans
			}
			for j := 0; j < 4; j++ {
				down := minus(cur, j)
				if !vis[down] {
					vis[down] = true
					q = append(q, down)
				}
				up := plus(cur, j)
				if !vis[up] {
					vis[up] = true
					q = append(q, up)
				}
			}

		}
		ans++
	}
	return -1
}

func minus(s string, i int) string {
	ret := ""
	for j, v := range s {
		if j == i {
			if v == '0' {
				ret += "9"
			} else {
				ret += string(v - 1)
			}
		} else {
			ret += string(v)
		}
	}
	return ret
}

func plus(s string, i int) string {
	ret := ""
	for j, v := range s {
		if j == i {
			if v == '9' {
				ret += "0"
			} else {
				ret += string(v + 1)
			}
		} else {
			ret += string(v)
		}
	}
	return ret
}
