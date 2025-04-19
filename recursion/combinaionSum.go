package recursion

import "fmt"

/*
38.
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
*/
func combinationSum(candidates []int, target int) [][]int {
	var result = make([][]int, 0)
	var arr = make([]int, 0, len(candidates))
	result = combination(result, candidates, arr, 0, target)
	return result
}

func combination(res [][]int, nums, arr []int, i, target int) [][]int {
	// base case
	if i >= len(nums) {
		return res
	}

	fmt.Println(arr, target, i)
	// result case
	if target == 0 {
		cp := make([]int, len(arr))
		copy(cp, arr)
		res = append(res, cp)
		return res
	}

	// pick
	if nums[i] <= target {
		arr = append(arr, nums[i])
		res = combination(res, nums, arr, i, target-nums[i])
		arr = arr[:len(arr)-1]
	}

	// not pick
	res = combination(res, nums, arr, i+1, target)

	return res
}
