package ac

func searchRange(nums []int, target int) []int {
    ansl := -1
    ansr := -1
    l, r := 0, len(nums)-1
    for l < r {
        mid := (l + r) / 2
        if nums[mid] == target {
            r = mid
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    if l >= 0 && l < len(nums) && nums[l] == target {
        ansl = l
    } else {
        return []int{-1, -1}
    }

    l, r = 0, len(nums)-1
    for l < r {
        mid := (l+r)/2 + 1
        if nums[mid] == target {
            l = mid
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    ansr = r

    return []int{ansl, ansr}
}
