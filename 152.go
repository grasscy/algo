package algo

func maxProduct(nums []int) int {
    max := nums[0]
    min := nums[0]
    pre := nums[0]
    for i := 1; i < len(nums); i++ {
        if pre*nums[i] > 0 {
            pre = pre * nums[i]
        }


    }
}
