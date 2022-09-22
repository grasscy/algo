package ac

func canFormArray(arr []int, pieces [][]int) bool {
	arri := 0
	for j := 0; j < len(pieces) && arri < len(arr); j++ {
		v := pieces[j]
		if v[0] == arr[arri] {
			arri++
			for i := 1; i < len(v); i++ {
				if v[i] != arr[arri] {
					return false
				}
				arri++
			}
			j = -1
		}
	}
	return arri == len(arr)
}
