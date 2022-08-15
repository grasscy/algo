package ac

var r [][]int
var used []int

func permute(nums []int) [][]int {
    r = [][]int{}
    used = make([]int, len(nums))
    ff(nums, []int{}, used)
    return r
}

func ff(nums []int, path []int, used []int) {
    if len(path) == len(nums) {
        r = append(r, append([]int{}, path...))
    }
    for i := 0; i < len(nums); i++ {
        if used[i] == 1 {
            continue
        }
        path = append(path, nums[i])
        used[i] = 1
        ff(nums, path, used)
        path = path[:len(path)-1]
        used[i] = 0
    }

}
