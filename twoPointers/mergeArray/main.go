package main

import "fmt"

func main() {
	mergeSortedArray([]int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, 3, 3)
}

func mergeSortedArray(nums1, nums2 []int, m, n int) {
	m, n, last := m-1, n-1, m+n-1
	//fmt.Println(nums1[m], nums2[n], nums1[k])
	for m+n > 0 {
		switch {
		case nums1[m] > nums2[n]:
			nums1[last] = nums1[m]
			m--
		case nums1[m] < nums2[n]:
			nums1[last] = nums2[n]
			n--
		case nums1[m] == nums2[n]:
			nums1[last] = nums2[n]
			n--
		}
		last--
	}
	fmt.Println(nums1)
}
