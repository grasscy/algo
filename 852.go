package algo

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid-1] > arr[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
