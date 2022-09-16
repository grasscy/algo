package ac

var temp []int

func sortArray(nums []int) []int {
	temp = make([]int, len(nums))
	sort(nums, 0, len(nums))
	return nums
}

func sort(nums []int, lo, hi int) {
	if lo == hi-1 {
		return
	}
	mid := (lo + hi) / 2
	sort(nums, lo, mid)
	sort(nums, mid, hi)
	merge(nums, lo, hi)
}

func merge(nums []int, lo, hi int) {
	for i := lo; i < hi; i++ {
		temp[i] = nums[i]
	}
	mid := (lo + hi) / 2
	i := lo
	j := mid
	for k := lo; k < hi; k++ {
		if i == mid {
			nums[k] = temp[j]
			j++
		} else if j == hi {
			nums[k] = temp[i]
			i++
		} else if temp[i] > temp[j] {
			nums[k] = temp[j]
			j++
		} else {
			nums[k] = temp[i]
			i++
		}
	}
}
