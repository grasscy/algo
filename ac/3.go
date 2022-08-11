package ac

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    m := map[uint8]int{}
    max, left := 0, 0
    for i := 0; i < len(s); i++ {
        v, ok := m[s[i]]
        if ok {
            if left < v+1 {
                left = v + 1
            }
        }
        m[s[i]] = i
        if max < i-left+1 {
            max = i - left + 1
        }
    }
    return max
}
