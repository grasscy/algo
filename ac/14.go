package ac

func longestCommonPrefix(strs []string) string {
    res := ""
    if len(strs) == 0 {
        return res
    }
    s := strs[0]
    r := 0
    for ; r < len(s); r++ {
        for j := 1; j < len(strs); j++ {
            if r >= len(strs[j]) {
                return res
            }
            if s[r] != strs[j][r] {
                return res
            }
        }
        res += string(s[r])
    }
    return res
}
