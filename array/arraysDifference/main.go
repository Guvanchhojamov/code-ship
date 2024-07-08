package main

import "fmt"

func main() {
	fmt.Println(twoArraysDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}

func twoArraysDifference(nums1, nums2 []int) [][]int {
	nums1Map, nums2Map := map[int]bool{}, map[int]bool{}
	answer := make([][]int, 2)
	for _, v := range nums1 {
		nums1Map[v] = true
	}

	for _, v := range nums2 {
		nums2Map[v] = true
	}

	for key := range nums1Map {
		if _, ok := nums2Map[key]; !ok {
			answer[0] = append(answer[0], key)
		}
	}

	for key := range nums2Map {
		if _, ok := nums1Map[key]; !ok {
			answer[1] = append(answer[1], key)
		}
	}
	return answer
}
