package daily

/*
1749. Maximum Absolute Sum of Any Subarray

You are given an integer array nums.
The absolute sum of a subarray [numsl, numsl+1, ..., numsr-1, numsr] is abs(numsl + numsl+1 + ... + numsr-1 + numsr).
Return the maximum absolute sum of any (possibly empty) subarray of nums.

Note that abs(x) is defined as follows:
If x is a negative integer, then abs(x) = -x.
If x is a non-negative integer, then abs(x) = x.


Example 1:
Input: nums = [1,-3,2,3,-4]
Output: 5
Explanation: The subarray [2,3] has absolute sum = abs(2+3) = abs(5) = 5.

Example 2:
Input: nums = [2,-5,1,-4,3,-2]
Output: 8
Explanation: The subarray [-5,1,-4] has absolute sum = abs(-5+1-4) = abs(-8) = 8.


Constraints:

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
*/
/*
  We need to find, subarray with maximum absolute sum:
     - nums[i] can be negative
     - we need abs sum
 1. Brute force:
    in 2 for loop
    sum elements and chec with maximum
    tc:O(n*n)
    sc: 1
 2. Can we solve this problem using prefix sum?
     What is the pattern?
     What if we need just maxSum not absolute sum?
     The absolute max sum == maxSum or minimum sum, eather of them

     For maximum  absolute sum we need to find minimumPrefixSum and maxPrefixSum and subtract them.
*/

func maxAbsoluteSum(nums []int) int {
	var prefixSum = 0
	var minPrefixSum = 0
	var maxPrefixSum = 0

	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		minPrefixSum = min(minPrefixSum, prefixSum)
		maxPrefixSum = max(maxPrefixSum, prefixSum)
	}

	return maxPrefixSum - minPrefixSum
}

/*
 [1,-3,2,3,-4]

 ps = -1
 mp = -2
 mxp = 3
  3 - (-2) = 5
 prefsSum = [1,-2,0,3,-1]

we need to find minimum prefix sum and maximum prefix sum from prefixSums;
and subsctract.
for example:
 i=1 j=4
 how ot find?
 psums[j] - psums[i-1]
 prefsSum[4] - prefsSum[0] =  -1 - 1 = -2
 [1,-3,2,3,-4]
 manual check: -3+2+3+(-4) = 2+(-4) = -2

*/
