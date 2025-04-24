package recursion

import "sort"

/*
Given a collection of candidate numbers (candidates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sum to target.
Each number in candidates may only be used once in the combination.
Note: The solution set must not contain duplicate combinations.


Example 1:
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

Example 2:
Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

Constraints:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/

/*
To ease avoid from duplicates, we need to sort the array, after sorting, all same elements, moving to nearest.
candidates = [2,5,2,1,2], target = 5
[1 2 2 2 5]
result = 1,2,2; 5
1,2,2 second 2
1,2,2 third 2   are duplicates, so we need ignore them.
  others not equal with sum, so ignore them too.
*/

/*
 How can we solve this problem?
 	1. Can we use brute force?
 how we can use pick and not pick method in there?
	what we need?
	arr - witch contains combinations
	res - contains answer
	i- index of element
	target - calculate sum, decreasing chosen value.
starting from each candidate we need to check: 	start-> n-1 all possibilities.
for idx= range n-1
	fn (res, nums,  arr, idx, i, target) []int:
		result case:
			target == 0:
				res = append.arr -> return
		base case:
			i == len(nums):
				return ans
		recursion case:
		// pick
if arr[i] <= target:    // how to avoid duplicates in there?  we need to check if nums[idx] != arr[idx-1] when i > idx  i must be bigger than
	index otherwise go to not pick.
 arr = append(arr[i])
 fn(res, nums, arr, i+1, target - arr[i]):
 arr = arr[len(arr)-1]   // remove
		// not pick
	fn()

*/

func combinationSum2(candidates []int, target int) [][]int {
	var res = make([][]int, 0, 0)
	sort.Ints(candidates)
	var arr = make([]int, 0, 0)
	res = fnnn(res, candidates, arr, 0, target)
	return res
}

func fnnn(res [][]int, nums, arr []int, idx, target int) [][]int {
	// answer case:
	if target == 0 {
		res = append(res, append([]int{}, arr...)) // to avoid mutation
		return res
	}
	// base case:
	if idx == len(nums) {
		return res
	}

	// pick, avoid duplicates
	for i := idx; i < len(nums); i++ {
		if i > idx && nums[i] == nums[i-1] { //why i>idx idx is stating point, and if elements same from start->n-1, before start point doesnt matter.
			continue
		}
		if nums[i] > target {
			break
		}
		arr = append(arr, nums[i])
		res = fnnn(res, nums, arr, i+1, target-nums[i])
		arr = arr[:len(arr)-1] // remove added
	}

	return res
}
