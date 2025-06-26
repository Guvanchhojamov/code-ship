package sliding_window

/*
1423. Maximum Points You Can Obtain from Cards

There are several cards arranged in a row,
and each card has an associated number of points.
The points are given in the integer array cardPoints.

In one step, you can take one card from the beginning or from the end of the row.
You have to take exactly k cards.

Your score is the sum of the points of the cards you have taken.

Given the integer array cardPoints and the integer k,
return the maximum score you can obtain.

Example 1:

Input: cardPoints = [1,2,3,4,5,6,1], k = 3
Output: 12
Explanation: After the first step, your score will always be 1. However, choosing the rightmost card first will maximize your total score. The optimal strategy is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 12.
Example 2:

Input: cardPoints = [2,2,2], k = 2
Output: 4
Explanation: Regardless of which two cards you take, your score will always be 4.
Example 3:

Input: cardPoints = [9,7,7,9,7,7,9], k = 7
Output: 55
Explanation: You have to take all the cards. Your score is the sum of points of all cards.


Constraints:

1 <= cardPoints.length <= 10^5
1 <= cardPoints[i] <= 10^4
1 <= k <= cardPoints.length
*/

/*
 pick from start k, back 0
 start k-1, back k-(k-1).
 start k-2 ,back k-(k-2)
 start k-i, back k-(k-i).
  while k-i>=0
   increase sum or decrease sum.

   edge case:
   if k == len(nums) return sum(nums).
   tc: O(n) overall in worth case.
   sc: O(1)
*/

func maxScore(nums []int, k int) int {
	var res = 0

	if k == len(nums) {
		for i := range nums {
			res += nums[i]
		}
		return res
	}

	var lsum, rsum = 0, 0
	for i := range k {
		lsum += nums[i]
	}
	var maxsum = lsum
	var ridx = len(nums) - 1
	for i := k - 1; i >= 0; i-- {
		lsum = lsum - nums[i]
		rsum = rsum + nums[ridx]
		ridx--
		maxsum = max(maxsum, lsum+rsum)
	}
	return maxsum
}
