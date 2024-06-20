package main

func main() {
	mergeSortedArray([]int{2, 5, 6, 0, 0, 0}, []int{1, 2, 3}, 3, 3)
}

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	m, n, last := m-1, n-1, m+n-1
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[last] = nums1[m]
			m--
		} else {
			nums1[last] = nums2[n]
			n--
		}
		last--
	}

	for n >= 0 {
		nums1[last] = nums2[n]
		n--
		last--
	}
}
