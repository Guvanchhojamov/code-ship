package greedy

/*
55. Jump Game
You are given an integer array nums.
You are initially positioned at the array's first index,
and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

Constraints:
1 <= nums.length <= 10^4
0 <= nums[i] <= 10^5

*/
/*
 We given nums, start is th 0-th position.
	- on each position we can jump x posiitions - x = a[i] - the value of this potion.
	- need to find can we reach the end of array or can not? return boolean.
	- out of the end border means we can reach end.
 - in the first example we can reach anyway, because we don't have 0 jumps.
 - but in the second if we start and go continuosly we may jump to 0 and stuck in 0.
	and wrongly return false.
	but we need to check other ways.
 - but sometimes we really cant reach the end, in that situtation we need return exactly faslse.
  the problem is can we reach from any way, or cannot reach from anyway.

1. approach:
	- start from 0 index, check for each valie.
	- start from the next value.
	do until reaches end, if reached return true and break.
	else return false.
tc: n^2.
sc: O(1).
2. Can we optimize?
	 arr = [2, 3, 1, 0, 4]
 in there we have 0 but we can reach end. if we jumt 1 in first instead of 2.
1. Keep track maxIndex on each step.
2. Check is prev index can reach to current index?
	If it is cannot then its impossible.
3. if current index are greater than maxindex then also impossible to reach end.
	because the maxIndex always must be after or equal the currentIndex, otherwise how
	jumper jumed to current? So it is impossible return false.
*/

func canJump(nums []int) bool {
	var maxJumpIndex = nums[0]

	for i, steps := range nums {
		if i > maxJumpIndex {
			return false
		}
		maxJumpIndex = max(maxJumpIndex, i+steps) // jumps
	}

	return true
}

/*
[2, 3, 1, 0, 4]
i=4
mx=8
 true

[3,2,1,0,4]
i=4
mx=3
false
*/
