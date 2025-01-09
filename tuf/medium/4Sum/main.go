package main

import "sort"

/*
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

Example 1:
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]

Constraints:
1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109
*/
/*

Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]

*/

/*
after sorted
[1,0,-1,0,-2,2]

-2 -1 0 0 1 2
*/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var k, l int
	var sum int
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		} // skip same values when moving i

		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			} // skip same values when moving the j
			k = j + 1
			l = len(nums) - 1
			for k < l {
				sum = nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < len(nums) && nums[k] == nums[k-1] {
						k++
					}
					for l > j && l != len(nums)-1 && nums[l] == nums[l+1] {
						l--
					}
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return ans
}
