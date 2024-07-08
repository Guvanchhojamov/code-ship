package main

import "fmt"

func main() {
	twoArraysDifference([]int{1, 2, 3}, []int{2, 4, 6})
}

func twoArraysDifference(nums1, nums2 []int) [][]int {
	nums1Map, nums2Map := map[int]bool{}, map[int]bool{}
	for _, v := range nums1 {
		nums1Map[v] = true
	}

	for _, v := range nums2 {
		nums2Map[v] = true
	}
	
	for i
	
	fmt.Println(nums1Map, nums2Map)
	return [][]int{}
}
