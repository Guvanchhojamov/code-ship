package greedy

import "math"

/*
Maximum Subarray
Given an integer array nums, find the subarray with the largest sum, and return its sum.
Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:
Input: nums = [1]
Output: 1

Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.


Constraints:

1 <= nums.length <= 10^5
-104 <= nums[i] <= 10^4


Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

we need to find subarray with max sum. subarray - continues part of some array.
bf:
 - generate all subarrays. in 2 loop
 - check with max, return max at the end.
tc: N*N
sc: 1
use prefix sum.
nums = [-2,1,-3,4,-1,2,1,-5,4]
1. since we need max sum we restore sum to 0 each time if it goes down to negative.
2. compare currSum with currSum+next. take max of them, and compare with max.
 cs=-2; mx=-2
 1>-2+1; no, cs=-2+1; mx=-1
 -3>1+

*/
// Kanadies algo
func maxSubArray(nums []int) int {
	var curr_sum = 0
	var mx = math.MinInt
	for _, v := range nums {
		curr_sum += v
		mx = max(mx, curr_sum)

		if curr_sum < 0 {
			curr_sum = 0
		}
	}
	return mx
}
