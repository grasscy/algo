package ac

func maxArea(height []int) int {
    i, j := 0, len(height)-1
    r := min(height[i], height[j]) * (j - i)
    for i != j {
        if height[i] >= height[j] {
            j--
        } else {
            i++
        }
        s := min(height[i], height[j]) * (j - i)
        if s > r {
            r = s
        }
    }
    return r
}
