package ac

func maxProduct(nums []int) int {
    maxf := nums[0]
    minf := nums[0]
    ret := nums[0]
    for i := 1; i < len(nums); i++ {
        mx, mn := maxf, minf
        maxf = max(mx*nums[i], max(mn*nums[i], nums[i]))
        minf = min(mx*nums[i], min(mn*nums[i], nums[i]))
        ret = max(maxf, ret)
    }
    return ret
}
