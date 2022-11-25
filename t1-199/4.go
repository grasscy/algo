package pass

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	left := (m + n + 1) / 2
	right := (m + n + 2) / 2
	kth := findKth(nums1, nums2, 0, 0, left)
	kth2 := findKth(nums1, nums2, 0, 0, right)
	return float64(kth+kth2) / 2.0
}

func findKth(nums1 []int, nums2 []int, i, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}
	midv1, midv2 := 0, 0
	if i+k/2-1 < len(nums1) {
		midv1 = nums1[i+k/2-1]
	} else {
		midv1 = math.MaxInt32
	}

	if j+k/2-1 < len(nums2) {
		midv2 = nums2[j+k/2-1]
	} else {
		midv2 = math.MaxInt32
	}

	if midv1 < midv2 {
		return findKth(nums1, nums2, i+k/2, j, k-k/2)
	} else {
		return findKth(nums1, nums2, i, j+k/2, k-k/2)
	}
}
