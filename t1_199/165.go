package t1_199

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	i := 0
	lmax := max(len(v1), len(v2))
	for ; i < lmax; i++ {
		var atoi1, atoi2 int
		if i < len(v1) {
			atoi1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			atoi2, _ = strconv.Atoi(v2[i])
		}
		if atoi1 < atoi2 {
			return -1
		} else if atoi1 > atoi2 {
			return 1
		}
	}

	return 0
}
