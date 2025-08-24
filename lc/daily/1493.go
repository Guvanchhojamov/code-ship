package daily

import "math"

/*
1493. Longest Subarray of 1's After Deleting One Element

Given a binary array nums, you should delete one element from it.
Return the size of the longest non-empty subarray containing only 1's in the resulting array.
Return 0 if there is no such subarray.

Example 1:

Input: nums = [1,1,0,1]
Output: 3

Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.

Example 2:
Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
Example 3:

Input: nums = [1,1,1]
Output: 2
Explanation: You must delete one element.


Constraints:

1 <= nums.length <= 10^5
nums[i] is either 0 or 1.

*/
/*
  We need finds lognest subarray after deletetinon 0 element.
  - must delete 1 element from subarray, if no 0 must delete 1.
  - subarray - part of continues of elements in array.
  - array is binary only 0,1.
  - range is big.
 1. generate al subarrays, remove 1 zero or 0. compare length with max.
	[0,1,1,1,0,1,1,0,1]
	 - how we can generate subarrays.
	   i= 0 -> n;
	     j = i -> n.
			if a[j] == 0 : z++
			if z == 2: break
			else count++
		mx = max(mx,count-1)
		z=0; count = 0;
    tc: n*n;
	sc: 1.
2. How can we optimze?
   For subarrays, can we use sliding window with 2 pointers?
	[0,1,1,1,0,1,1,0,1].
	i,j = 0, 0. z = 0.
	approach:
	- z - to keep zeros count.
	- j-i+1 - window size to compare with prev max.
	- start i,j starting point.
	- until i,j < n move window.
	   if a[j] == 0: z ++
	   while z>1:
		  if a[i] == 0: z--
		  i++
	   mx = max(mx, j-i+1)
	  j++
	- edge case: what is no zeros in array? we must delete 1 elemnt any. result is: len(arr)-1.
	- what if arr is all elements are 0. we need return 0.
	  We can check edge cases before processing array.  This will take N time.
	  Or can we do it in processing?


	tc: in the worth case we iiterate over array 2 times. n for i and n for j so, n+n = 2n = N.
	sc: 1.
*/
func longestSubarray(nums []int) int {
	var mx = math.MinInt
	var i, j, zeros = 0, 0, 0
	for i < len(nums) && j < len(nums) {
		if nums[j] == 0 {
			zeros++
		}
		for zeros > 1 {
			if nums[i] == 0 {
				zeros--
			}
			i++
		}
		mx = max(mx, j-i) // we need to remove -1 element too, so this is why dont do j-i-1. just j-i.
		j++
	}
	return mx
}

/*
 after editioral, we know we dont need to check ones count, the logic covers all ones. edge case.
 because we use j-i = it is endi-starti = this is len(nums)-1.
 But how he cover all zeros array?  [0 0 0 0] - because at the end starti == endi so j-i=. we return 0.
*/
