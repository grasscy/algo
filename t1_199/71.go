package t1_199

import "strings"

func simplifyPath(path string) string {
	split := strings.Split(path, "/")
	ans := []string{}
	for _, v := range split {
		if v == "" {
		} else if v == "." {

		} else if v == ".." {
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		} else {
			ans = append(ans, v)
		}
	}
	return "/" + strings.Join(ans, "/")
}
