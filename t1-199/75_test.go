package pass

import "testing"

func Test_sortColors(t *testing.T) {
	sortColors([]int{1, 2, 0})
	sortColors([]int{2, 0, 1})
	sortColors([]int{2, 0, 2, 1, 1, 0})

}
