package main

/*
496. Leetcode

nums1 = [4,1,2],
nums2 = [1,3,4,2]

	out:   [-1,3,-1]
*/
func main() {
	nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	numsMap := map[int]int{nums2[len(nums2)-1]: -1}
	result := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		if len(stack) > 0 {
			switch {
			case nums2[i] < stack[len(stack)-1]:
				numsMap[nums2[i]] = stack[len(stack)-1]
			case nums2[i] > stack[len(stack)-1]:
				for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				}
				if len(stack) > 0 {
					numsMap[nums2[i]] = stack[len(stack)-1]
				} else {
					numsMap[nums2[i]] = -1
				}
			}
		}
		stack = append(stack, nums2[i])
	}

	for _, v := range nums1 {
		result = append(result, numsMap[v])
	}
	return result
}
