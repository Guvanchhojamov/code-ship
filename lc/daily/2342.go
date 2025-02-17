package daily

import "sort"

// 2342. Max Sum of a Pair With Equal Sum of Digits
/*
You are given a 0-indexed array nums consisting of positive integers.
You can choose two indices i and j, such that i != j,  and the sum of digits of the number nums[i] is equal to that of nums[j].

Return the maximum value of nums[i] + nums[j] that you can obtain over all possible indices i and j that satisfy the conditions.



Example 1:

Input: nums = [18,43,36,13,7]
Output: 54
Explanation: The pairs (i, j) that satisfy the conditions are:
- (0, 2), both numbers have a sum of digits equal to 9, and their sum is 18 + 36 = 54.
- (1, 4), both numbers have a sum of digits equal to 7, and their sum is 43 + 7 = 50.
So the maximum sum that we can obtain is 54.

Example 2:
Input: nums = [10,12,19,14]
Output: -1
Explanation: There are no two numbers that satisfy the conditions, so we return -1.

Constraints:
1 <= nums.length <= 105
1 <= nums[i] <= 109

*/

/*
  What we need?
    Find max Sum from elements where:
      - sumDigits(A) == sumDigits(B)
      - sum(A,B).

    Key points:
        - We need nums just where digts sum are equeal.
        - So all nums with same digits sums creates same key. (We can use hash map).
        - sumsMap:
            key = number digits sum.
            value = numbers sum (with same digits sum).
        - return max(sumsMap.Values)
        - If len(map) == len(arr) then we can sure tan any number has own different digits sum .
            No numbers with same digits sum.
            So we return -1
        And we return maximum value in map.
        - We need helper function to get Numbers digits sum.
 Some test cases:
 nums = [18,43,36,13,7]
  9:18+36
  7:43+7
  4:13
  len(map) != len(arr) so we have nums with same diigt nums.
  return max = 18+36=54
  tc: O(N)
  sc:O(n)
     is it ok? Yes.
  Follow up: Can we optimize space complexity for this problem.
*/

/*
How about helper function?
n=153

	while n>0:
	  s+=n % 10
	  n=n/10
	return s

0+3+5+1 = 9

	n*logN
*/
func maximumSum(nums []int) int {
	var sumsMap = make(map[int][]int, 0)
	var maxSum = 0 // since a[i] minimum constraints starts with 1.

	for i := 0; i < len(nums); i++ {
		digsSum := digitsSum(nums[i])
		sumsMap[digsSum] = append(sumsMap[digsSum], nums[i])
	}
	if len(sumsMap) == len(nums) {
		return -1
	}

	for _, arr := range sumsMap {
		sort.Ints(arr)
		if len(arr) >= 2 {
			maxSum = max(maxSum, arr[len(arr)-1]+arr[len(arr)-2])
		}
	}

	return maxSum
}

func digitsSum(n int) int {
	var sum = 0
	for n > 0 {
		sum += int(n % 10)
		n /= 10
	}
	return sum
}

/*
 Solved using tips.
   What learned?
    - Learn HEAP usage for this related questions.
    - n % 10 -> always gives float val, convert to int if needed.
    - Focus on small parts of code.
    - !important, Map range usage different from get map value usage.
*/
/*
 nums = [18,43,36,13,7]
 maxSum = 0
 9:18
 7:43
 0 > 18+36 = 54, maxSum = 54.
 4:13
 54, 43+7 = 54,50 = 54
 3!=5
 return 54.
*/
/*
[420,464,491,218,439,153,482,169,411,93,147,50,347,210,251,366,401]

map[3:210 5:451 6:831 8:251 9:153 12:240  14:1784 15:366 16:608  ]
13:229
20:398
17:269
11:535
*/
