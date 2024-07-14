package main

import "fmt"

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	intersectionII([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
}

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]
// Explanation: [4,9] is also accepted.

func intersection(nums1 []int, nums2 []int) []int {
	nums1Map, nums2Map := map[int]bool{}, map[int]bool{}
	var result []int
	for _, v := range nums1 {
		nums1Map[v] = true
	}
	fmt.Println(nums1Map)
	for _, v := range nums2 {
		if _, ok := nums2Map[v]; !ok {
			nums2Map[v] = true
			if _, exists := nums1Map[v]; exists {
				result = append(result, v)
			}
		}
	}

	return result
}

func intersectionII(nums1 []int, nums2 []int) []int {
	nums1Map, nums2Map := map[int]bool{}, map[int]bool{}
	var result []int
	for _, v := range nums1 {
		nums1Map[v] = true
	}
	fmt.Println(nums1Map)
	for _, v := range nums2 {
		nums2Map[v] = true
	}

	return result
}
