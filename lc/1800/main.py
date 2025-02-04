/*
1800. Maximum Ascending Subarray Sum
Given an array of positive integers nums, return the maximum possible sum of an ascending subarray in nums.
A subarray is defined as a contiguous sequence of numbers in an array.
A subarray [numsl, numsl+1, ..., numsr-1, numsr] is ascending if for all i where l <= i < r, numsi  < numsi+1. Note that a subarray of size 1 is ascending.

Example 1:
Input: nums = [10,20,30,5,10,50]
Output: 65
Explanation: [5,10,50] is the ascending subarray with the maximum sum of 65.

Example 2:
Input: nums = [10,20,30,40,50]
Output: 150
Explanation: [10,20,30,40,50] is the ascending subarray with the maximum sum of 150.

Example 3:
Input: nums = [12,17,15,13,10,11,12]
Output: 33
Explanation: [10,11,12] is the ascending subarray with the maximum sum of 33.


Constraints:

1 <= nums.length <= 100
1 <= nums[i] <= 100
*/

class Solution:
    def maxAscendingSum(self, nums: List[int]) -> int:
        max_sum = nums[0]
        curr_sum = nums[0]

        for val in nums:
            max_sum = max(max_sum, val)

        for i in range(1,len(nums)):
            if nums[i-1] < nums[i]:
                curr_sum+=nums[i]
            else:
                max_sum = max(max_sum, curr_sum)
                curr_sum = nums[i]
        return max(max_sum,curr_sum)

/*
  [10,20,30,5,10,50]
  cs:
  ms: 65

*/

/*
  What we need ?
   Need to find maxSum of increasing subarray.
   Since, constraints is small we can use for this taks brute foce solution first, then we can optimize.
   Key points:
    - one element a[i] is individually increasing subarray.
    - If no any increasing subbaray in given array we need return maximum element in array.

   1. Brute force.
     for i in nums-1:
        if a[i] < a[i+1]:
            sum+=a[i];
        else:
            maxSum = max(maxSum, sum)
            sum = a[i]
    maxSum = max(maxSum, sum)
    return maxSum
*/