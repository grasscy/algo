package t1_199

func removeElement(nums []int, val int) int {
	fast, slow := 0, 0

	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
