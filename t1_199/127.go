package t1_199

func ladderLength(beginWord string, endWord string, wordList []string) int {
    ans := 0
    found := false
    step := 1

    wordset := map[string]bool{}
    for _, v := range wordList {
        wordset[v] = true
    }
    q := []string{beginWord}
    for len(q) != 0 {
        step++
        size := len(q)
        for i := 0; i < size; i++ {
            t := q[0]
            q = q[1:]
            for j := 0; j < len(beginWord); j++ {
                for c := 'a'; c <= 'z'; c++ {
                    var nw string
                    if j+1 < len(beginWord) {
                        nw = t[:j] + string(c) + t[j+1:]
                    } else {
                        nw = t[:j] + string(c)
                    }
                    if !wordset[nw] {
                        continue
                    }
                    q = append(q, nw)
                    delete(wordset, nw)
                    if endWord == nw {
                        found = true
                        break
                    }
                }
            }
            if found {
                return step
            }
        }

    }

    return ans
}
