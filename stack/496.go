package stack

/*
 Leetcode 496. Next greater element I.
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var res = make([]int, 0, len(nums2))
	var st = make([]int, 0, len(nums2))
	var nextMap = make(map[int]int, 0)

	nextMap[nums2[len(nums2)-1]] = -1
	st = append(st, nums2[len(nums2)-1])

	for i := len(nums2) - 2; i >= 0; i-- {
		if nums2[i] < st[len(st)-1] {
			nextMap[nums2[i]] = st[len(st)-1]
			st = append(st, nums2[i])
		} else if len(st) == 0 {
			nextMap[nums2[i]] = -1
			st = append(st, nums2[i])
		} else {
			for len(st) > 0 && nums2[i] > st[len(st)-1] {
				st = st[:len(st)-1]
			}
			if len(st) == 0 {
				nextMap[nums2[i]] = -1
				st = append(st, nums2[i])
			} else {
				nextMap[nums2[i]] = st[len(st)-1]
				st = append(st, nums2[i])
			}
		}
	}

	for _, v := range nums1 {
		res = append(res, nextMap[v])
	}
	return res
}
