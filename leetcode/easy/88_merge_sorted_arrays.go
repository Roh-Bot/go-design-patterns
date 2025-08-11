package easy

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1ActualLengthIter := m + n - 1
	m = m - 1
	n = n - 1

	for m >= 0 && n >= 0 {
		if nums1[m] < nums2[n] {
			nums1[nums1ActualLengthIter] = nums2[n]
			n--
		} else {
			nums1[nums1ActualLengthIter] = nums1[m]
			m--
		}
		nums1ActualLengthIter--
	}

	// Copy remaining elements from nums2 if any
	for n >= 0 {
		nums1[nums1ActualLengthIter] = nums2[n]
		n--
		nums1ActualLengthIter--
	}
}
