package ac

func minElements(nums []int, limit int, goal int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    sub := goal - sum
    if sub < 0 {
        sub = -sub
    }
    ans := 0

    n := sub / limit
    ans += n
    if sub%limit!=0{
        ans++
    }

    return ans
}
