package t200_399

import (
	"strconv"
)

func isAdditiveNumber(num string) bool {
	return b(num, 0, []int{})
}

func b(num string, index int, tmp []int) bool {
	if index == len(num) {
		return len(tmp) >= 3
	}

	for i := index; i < len(num); i++ {
		if num[index] == '0' && i > index {
			return false
		}
		n, _ := strconv.Atoi(num[index : i+1])
		if len(tmp) >= 2 {
			if tmp[len(tmp)-1]+tmp[len(tmp)-2] != n {
				continue
			}
		}
		tmp = append(tmp, n)
		if b(num, i+1, tmp) {
			return true
		}
		tmp = tmp[:len(tmp)-1]
	}
	return false
}
