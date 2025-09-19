package greedy

/*
45. Jump Game II
You are given a 0-indexed array of integers nums of length n.
You are initially positioned at index 0.

Each element nums[i] represents the maximum length of a forward jump from index i.
In other words, if you are at index i, you can jump to any index (i + j) where:

0 <= j <= nums[i] and i + j < n
Return the minimum number of jumps to reach index n - 1.
The test cases are generated such that you can reach index n - 1.



Example 1:

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2.
Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [2,3,0,1,4]
Output: 2

Constraints:

1 <= nums.length <= 10^4
0 <= nums[i] <= 1000
It's guaranteed that you can reach nums[n - 1].

*/
/*
 We need to find minimum jumps count to reach the end of array.
 - the end is always reachable, its guranteed.
 - the jump steps asuume j.
	0 <= j <= nums[i]
	and i+j always < n

 [2,3,1,1,4]
  to find the minumum way.
  we need to go greedy and we need to choose best option in current time, in each step.
  - on start we can jump to n positions.
	our goal is reach end as soon as possible.
	so we need jump on each step at max.
  - 0 - 3,1 - take max = 3. jump to 1.
  - 2 - 1,1,4 - take max = 4. jump 3 step.
  - and we reach the end. in 2 jumps.

approach:
- iterate over nums:
   while j=i; j<i+a[i]:
	jumpi = max(jumpi,)
*/

func jumpGameII(nums []int) int {
	var currBundary, farest = 0, 0
	var jumps = 0
	for i := 0; i <= len(nums)-2; i++ {
		farest = max(farest, i+nums[i])
		if i == currBundary {
			currBundary = farest
			jumps++
		}
	}
	return jumps
}
