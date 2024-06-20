package main

func main() {
	mergeSortedArray([]int{2, 5, 6, 0, 0, 0}, []int{1, 2, 3}, 3, 3)
}

func mergeSortedArray(nums1, nums2 []int, m, n int) {
	i, j, last := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[last] = nums1[i]
			i--
		} else {
			nums1[last] = nums2[j]
			j--
		}
		last--
	}
	for j >= 0 {
		nums1[last] = nums2[j]
		j--
		last--
	}
}
