package ac

var ans [][]int
var visited map[int]bool

func permute(nums []int) [][]int {
	ans = [][]int{}
	visited = map[int]bool{}
	b(nums, []int{})
	return ans
}

func b(nums []int, tmp []int) {
	if len(tmp) == len(nums) {
		ans = append(ans, append([]int{}, tmp...))
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		tmp = append(tmp, nums[i])
		b(nums, tmp)
		tmp = tmp[:len(tmp)-1]
		visited[i] = false
	}
}
