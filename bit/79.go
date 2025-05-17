package bit

/*
78. Subsets
Given an integer array nums of unique elements,
return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

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
 We already solved this question, and
 we can solve this using recursion approach (not sure, need repeat )) )
 Lets try to solve with binary method.
    What is the response length?
		it is 2^len(nums)-1 right?
	1,2,3 :
[], [1], [2], [3],
[1,2], [1,3], [2,3] [1,2,3]
  so we need to generate bin 0->2^n-1
 000
 001
 010
 100
 101
 110
 111
 2^n == 1 << n
 subs=(1 << n) - 1
 for i=0 -> subs;
	for j:=0 -> len(nums)-1:
	0->8
     0->2
     if 0[j] == set bit:    j-> 0,1,2
 		list.append(nums[j])
	res.append(list)
*/
func subsets(nums []int) [][]int {
	var n = len(nums)
	var subs = (1 << n) - 1 // 2^n
	var res = make([][]int, 0, subs)

	for num := 0; num <= subs; num++ {
		list := []int{}
		for i := 0; i < n; i++ {
			if num&(1<<i) > 0 { // check ith count set or not set
				list = append(list, nums[i])
			}
		}
		res = append(res, list)
	}
	return res
}

/*
 000
 100 = 0
 001
 001
*/
