package ac

var r [][]int

func subsets(nums []int) [][]int {
    r = [][]int{}
    f(nums, 0, []int{})
    return r
}

func f(nums []int, index int, set []int) {
    r = append(r, append([]int{}, set...))
    for i := index; i < len(nums); i++ {
        set = append(set, nums[i])
        f(nums, i+1, set)
        set = set[:len(set)-1]
    }
}
