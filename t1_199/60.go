package t1_199

import "strconv"

var vis []bool
var fs []int
var ans string

func getPermutation(n int, k int) string {
	vis = make([]bool, n+1)
	fs = []int{1}
	for i := 1; i <= n; i++ {
		fs = append(fs, fs[i-1]*i)
	}

	b(n, k, 0, "")
	return ans
}

func b(n, k int, index int, tmp string) {
	if index == n {
		ans = tmp
		return
	}
	cnt := fs[n-1-index]
	for i := 1; i <= n; i++ {
		if vis[i] {
			continue
		}
		if cnt < k {
			k -= cnt
			continue
		}
		tmp += strconv.Itoa(i)
		vis[i] = true
		b(n, k, index+1, tmp)
		return
	}
}

//
//var ans string
//var vis []bool
//var cur int
//
//func getPermutation(n int, k int) string {
//    vis = make([]bool, n+1)
//    cur = 0
//    b(n, k, "")
//    return ans
//}
//
//func b(n int, k int, tmp string) {
//    if len(tmp) == n {
//        cur++
//        if k == cur {
//            ans = tmp
//        }
//        return
//    }
//    if k < cur {
//        return
//    }
//    for i := 1; i <= n; i++ {
//        if vis[i] {
//            continue
//        }
//        vis[i] = true
//        b(n, k, tmp+strconv.Itoa(i))
//        vis[i] = false
//    }
//}
