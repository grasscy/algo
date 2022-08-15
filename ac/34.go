package ac

func searchRange(nums []int, target int) []int {
    //leftmost := sort.SearchInts(nums, target)
    //if leftmost == len(nums) || nums[leftmost] != target {
    //    return []int{-1, -1}
    //}
    //rightmost := sort.SearchInts(nums, target + 1) - 1
    //return []int{leftmost, rightmost}

    lo, hi, mid := 0, len(nums)-1, 0
    r := -1
    for lo <= hi {
        mid = lo + (hi-lo)/2
        if nums[mid] == target {
            r = mid
            break
        }
        if nums[mid] > target {
            hi = mid - 1
        } else {
            lo = mid + 1
        }

    }
    if r == -1 {
        return []int{-1, -1}
    }
    i, j := r, r
    for nums[i] == target {
        i--
        if i < 0 {
            break
        }
    }
    for nums[j] == target {
        j++
        if j >= len(nums) {
            break
        }
    }
    return []int{i + 1, j - 1}
}
