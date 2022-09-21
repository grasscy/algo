package ac

var used map[int]bool
var ans [][]int

func permute(nums []int) [][]int {
	ans = [][]int{}
	used = map[int]bool{}
	backtrack(nums, 0, []int{})
	return ans
}

func backtrack(nums []int, n int, path []int) {
	if n == len(nums) {
		ans = append(ans, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[nums[i]] {
			continue
		}
		used[nums[i]] = true
		path = append(path, nums[i])
		backtrack(nums, n+1, path)
		used[nums[i]] = false
		path = path[:len(path)-1]
	}
}
