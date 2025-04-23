package recursion

import "sort"

/*
Given an integer array nums that may contain duplicates,
return all possible subsets (the power set).

The solution set must not contain duplicate subsets.
Return the solution in any order.

Example 1:
Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

Example 2:
Input: nums = [0]
Output: [[],[0]]


Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
*/

/*
to avoid duplicated 1st we need to sort the given arr,
so the same numbers should be in the same place.
1. Sort nums
2. do the pick and not pick subset logic, but we can change something.
  we need to check where i>startIndex && a[i] != a[i-1]
	so after start point we take only unique elements in there, and skip all same elements, or
we take each element only once, skipping other same elements.
*/
/*
  base case:
	i == len(nums)
	return res.append(arr)

	// pick
  for i:=currIdx; i<len(nums);i++:
	if i > currIdx && nums[i] != nums[i-1]:
		append(arr,nums[i])
		fn(res,nums,arr,i,currIdx)
		arr = arr[:len(arr)-1]
*/

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums) // to handle duplicates
	var res [][]int
	var curr []int
	backtrack(nums, 0, curr, &res)
	return res
}

func backtrack(nums []int, start int, curr []int, res *[][]int) {
	*res = append(*res, append([]int{}, curr...)) // Add a copy of current subset

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue // skip
		}
		curr = append(curr, nums[i])
		backtrack(nums, i+1, curr, res)
		curr = curr[:len(curr)-1] // Backtrack - remove pciked element
	}
}
