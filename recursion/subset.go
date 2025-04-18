package recursion

/*
78. Subsets
Given an integer array nums of unique elements, return all possible subsets (the power set).
The solution set must not contain duplicate subsets.
Return the solution in any order.


Example 1:
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
Input: nums = [0]
Output: [[],[0]]


Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.
*/

/*
 We need use there pick and not pick method for, generate all subsequences.
starting from [] arr,
	pick and print or
	not pick and print
Or
	we can use pick and not pick in other order.
start from full arr
second starts from empty array.
*/

func subsets(nums []int) [][]int {
	result := [][]int{}
	return sub(result, nums, []int{}, 0)
}

func sub(res [][]int, nums []int, arr []int, i int) [][]int {
	if i == len(nums) {
		// Why we copy in there, because of mutation. in arrays. Should need to learn what is the mutation.
		cp := make([]int, len(arr))
		copy(cp, arr)
		res = append(res, cp)
		return res
	}

	// not pick case
	res = sub(res, nums, arr, i+1)

	// pick case
	arr = append(arr, nums[i])
	res = sub(res, nums, arr, i+1)

	return res
}
