package recursion

/*
Given an array nums of distinct integers, return all the possible permutations.
You can return the answer in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:
Input: nums = [1]
Output: [[1]]

Constraints:
1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/

/*
  What is the permutation?
  Permutation is all possible combinations,
the max combinations for n is n!.
     1. brute force.
 	start from 1
	create seenMap  and add picked elements to this seenMap.
	each time start from 0 and check if it si in the seen map  then continue else, take and
	start again from 0 -> n  until curr.len == nums.len.
	start from each element. and do this.
	  when backtrack dont forget remove picked number from map and current array.
tc: for each num we check all possible combinations, so, n*n!
sc: 2N = N

For combinations and subsets the start will increase each time.
But for permutations we need to take every element, just replacing the places.
For replace the place we each time start from different position, for ex: start, start+1,....
*/

func permute(nums []int) [][]int {
	var res = make([][]int, 0)
	var seen = make(map[int]bool, 0)
	var curr = make([]int, 0)
	fnPer(&res, nums, curr, seen)
	return res
}

func fnPer(res *[][]int, nums, curr []int, seen map[int]bool) {
	// base case
	if len(curr) == len(nums) {
		*res = append(*res, append([]int{}, curr...))
		return
	}

	// recursive case:
	for i := 0; i < len(nums); i++ {
		if seen[nums[i]] {
			continue
		}
		seen[nums[i]] = true
		curr = append(curr, nums[i])
		fnPer(res, nums, curr, seen)
		curr = curr[:len(curr)-1]
		seen[nums[i]] = false
	}
	return
}

/*
 What we can optimize?
  We can optimize in there the space, we can not use the hashmap to store elements.
 We can just swap:
 	how we can swap then?
  start from start until end swap and when reached out end, add it to answer,
 when backtracked swap back to get starting state.
*/

func permute2(nums []int) [][]int {
	var res = make([][]int, 0)
	// var curr = make([]int, 0)

	fnPer2(&res, nums, 0)

	return res
}

func fnPer2(res *[][]int, nums []int, start int) {
	// base case
	if start == len(nums) {
		*res = append(*res, append([]int{}, nums...))
		return
	}

	// recursive case:
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]

		fnPer2(res, nums, start+1)

		nums[start], nums[i] = nums[i], nums[start]
	}
	return
}
