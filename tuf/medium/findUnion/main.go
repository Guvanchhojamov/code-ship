package main

import "fmt"

func main() {
	nums1 := []int{1, 1, 2, 3, 4, 5}
	nums2 := []int{2, 3, 4, 4, 5, 6, 6, 6}
	findUnion(nums1, nums2)
}

/*
1 1 2 3 4 5
2 3 4 4 5 6

out: 1 2 3 4 5 6
1 2 3 4 5 6
*/
func findUnion(nums1, nums2 []int) {
	i, j := 0, 0
	result := []int{}
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			if (len(result) > 0 && result[len(result)-1] != nums1[i]) || len(result) == 0 {
				result = append(result, nums1[i])
			}
			i++
		} else {
			if (len(result) > 0 && result[len(result)-1] != nums2[j]) || len(result) == 0 {
				result = append(result, nums2[j])
			}
			j++
		}
	}

	for i < len(nums1) {
		if (len(result) > 0 && result[len(result)-1] != nums1[i]) || len(result) == 0 {
			result = append(result, nums1[i])
		}
		i++
	}

	for j < len(nums2) {
		if (len(result) > 0 && result[len(result)-1] != nums2[j]) || len(result) == 0 {
			result = append(result, nums2[j])
		}
		j++
	}

	fmt.Println(result)
}
