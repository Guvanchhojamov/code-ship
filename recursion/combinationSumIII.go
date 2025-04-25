package recursion

/*
Find all valid combinations of k numbers that
sum up to n such that the following conditions are true:
	Only numbers 1 through 9 are used.
	Each number is used at most once.
Return a list of all possible valid combinations.
The list must not contain the same combination twice, and the combinations may be returned in any order.

Example 1:
Input: k = 3, n = 7
Output: [[1,2,4]]

Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.
Example 2:
Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Explanation:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.

Example 3:
Input: k = 4, n = 1
Output: []
Explanation: There are no valid combinations.
Using 4 different numbers in the range [1,9],
the smallest sum we can get is 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.
Constraints:
2 <= k <= 9
1 <= n <= 60

*/

/*
  What we need
 1. Find unique combinations for nums 1..9 for length k
 2. Calculate sum of these combinations.
 3. If result equal to n then add it to answer.
*/

func combinationSum3(k int, n int) [][]int {
	var res = make([][]int, 0)
	var current = make([]int, 0)
	var sum = 0
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fnCombs(&res, nums, current, 0, k, sum, n)

	return res
}

func fnCombs(res *[][]int, nums, current []int, start, k, sum, target int) {
	// base case:
	if k == 0 && sum == target {
		*res = append(*res, append([]int{}, current...)) // to avoid array mutation we need create new array.
		return
	}

	if start == len(nums) {
		if k == 0 && sum == target {
			*res = append(*res, append([]int{}, current...))
		}
		return
	}

	// recursion case
	for i := start; i < len(nums); i++ {
		if nums[i] > target || sum > target || (i > start && nums[i] == nums[i-1]) { // we can simplify in there?
			continue
		}
		current = append(current, nums[i])
		fnCombs(res, nums, current, i+1, k-1, sum+nums[i], target)
		current = current[:len(current)-1]
	}
	return
}

// Optimized after chat gpt review

func fnCombsOpt(res *[][]int, nums, current []int, start, k, sum, target int) {
	if k == 0 && sum == target {
		*res = append(*res, append([]int{}, current...))
		return
	}
	if k < 0 || sum > target {
		return
	}

	for i := start; i < len(nums); i++ {
		if sum+nums[i] > target {
			continue
		}
		current = append(current, nums[i])
		fnCombs(res, nums, current, i+1, k-1, sum+nums[i], target)
		current = current[:len(current)-1]
	}
}
