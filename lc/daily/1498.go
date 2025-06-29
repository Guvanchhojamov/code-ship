package daily

/*

1498. Number of Subsequences That Satisfy the Given Sum Condition

You are given an array of integers nums and an integer target.
Return the number of non-empty subsequences of nums
such that the sum of the minimum and maximum element on it is less or equal to target.
Since the answer may be too large, return it modulo 109 + 7.



Example 1:

Input: nums = [3,5,6,7], target = 9
Output: 4
Explanation: There are 4 subsequences that satisfy the condition.
[3] -> Min value + max value <= target (3 + 3 <= 9)
[3,5] -> (3 + 5 <= 9)
[3,5,6] -> (3 + 6 <= 9)
[3,6] -> (3 + 6 <= 9)

Example 2:
Input: nums = [3,3,6,8], target = 10
Output: 6
Explanation: There are 6 subsequences that satisfy the condition. (nums can have repeated numbers).
[3] , [3] , [3,3], [3,6] , [3,6] , [3,3,6]

Example 3:
Input: nums = [2,3,3,4,6,7], target = 12
Output: 61
Explanation: There are 63 non-empty subsequences, two of them do not satisfy the condition ([6,7], [7]).
Number of valid subsequences (63 - 2 = 61).


Constraints:

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^6
1 <= target <= 10^6

*/

/*
- number of subsequences with min+max<=target.
 brute force:
	Generate all subsecuences and compare sum.
	return count.
- generate with backtrack pick non pick. O(n)
- find min and max on generating. sum on base case and compare with k.

tc: N
sc: N for stack. N for tmp array. 2N.
since count myght be large awe myght got TLE error. but myght not.

Any mathematical or other solution?
- find for each num, how much subarrays he is min?
- how much he is max?
how to know he is min+max <= target.
Input: nums = [2,3,3,4,6,7], target = 12
Output: 61
1. Sort nums since subsequence we can sort nums.
2. after sorting the all mins 0 to n/2; maxes- n/2 to n.
  [3,5,6,7]

Ok. lets try first approach.

1. sort array.
 nums;
 arr - tmp arr.
 mxSum - min, max sum of arr.
 i - ith element.
	fn(nums, arr)
 base case:
	sum > target or i=n
 recursion case:
   not pick
	fn(nums,arr,mxSum,&res, i+1)
   pick
   arr= append(nums[i])
	fn(nums,arr,mxSum,&res, i+1)

	return res.
	mod 1o^7 res each time.

*/


func numSubseq(nums []int, target int) int {
    
}
