package main

func main() {
	intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
}

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]
// Explanation: [4,9] is also accepted.

func intersection(nums1 []int, nums2 []int) []int {
	nums1Map, nums2Map := map[int]bool{}, map[int]bool{}
	for _, v := range nums1 {
		nums1Map[v] = true
	}
	for _, v := range nums2 {
		nums2Map[v] = true
	}

	var result []int
	for key := range nums1Map {
		if _, ok := nums2Map[key]; ok {
			result = append(result, key)
		}
	}
	return result
}
