package t200_399

var temp []Pair
var count []int

type Pair struct {
	Val int
	Id  int
}

func countSmaller(nums []int) []int {
	temp = make([]Pair, len(nums))
	count = make([]int, len(nums))

	arr := []Pair{}
	for i, v := range nums {
		arr = append(arr, Pair{v, i})
	}
	sort(arr, 0, len(nums))
	return count
}

func sort(nums []Pair, lo, hi int) {
	if lo == hi-1 {
		return
	}
	mid := (lo + hi) / 2
	sort(nums, lo, mid)
	sort(nums, mid, hi)
	merge(nums, lo, hi)
}

func merge(nums []Pair, lo, hi int) {
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
			count[nums[k].Id] += j - mid
		} else if temp[i].Val > temp[j].Val {
			nums[k] = temp[j]
			j++
		} else {
			nums[k] = temp[i]
			i++
			count[nums[k].Id] += j - mid
		}
	}
}
