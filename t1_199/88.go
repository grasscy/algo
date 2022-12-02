package t1_199

func merge(nums1 []int, m int, nums2 []int, n int) {
	p := len(nums1) - 1
	for m >= 1 && n >= 1 {
		if nums1[m-1] >= nums2[n-1] {
			nums1[p] = nums1[m-1]
			m--
		} else {
			nums1[p] = nums2[n-1]
			n--
		}
		p--
	}
	for m >= 1 {
		nums1[p] = nums1[m-1]
		m--
		p--
	}
	for n >= 1 {
		nums1[p] = nums2[n-1]
		n--
		p--
	}
}
