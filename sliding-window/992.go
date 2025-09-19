package sliding_window

/*

Test Result
992. Subarrays with K Different Integers

Given an integer array nums and an integer k, return the number of good subarrays of nums.

A good array is an array where the number of different integers in that array is exactly k.

For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.
A subarray is a contiguous part of an array.

Example 1:

Input: nums = [1,2,1,2,3], k = 2
Output: 7
Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2]
Example 2:

Input: nums = [1,2,1,3,4], k = 3
Output: 3
Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].


Constraints:

1 <= nums.length <= 2 * 104
1 <= nums[i], k <= nums.length

*/
/*
 we need to return exactly k diff integers.
 brute force:
 - generate all subarrays.
 - store integers in map
 when len map ==k inc count.
 tc: n*n
 sc: O(k)
 How cnan we optimze this solution?
 Can we use sliding window?
    Input: nums = [1,2,1,2,3], k = 2
 one way use math approahc x=count<=k and y = count <=k-1
 res = x-y
  tc: n+n
  sc: n+n
*/
/*
 How can we optimze it using sliding window?
 This is typycal two pointer approach not work there,
 because we miss some subarrays.
 Ok, lets do previously known approach
 x: nums <=k; and y: nums <=k-1
 res = x-y. This works always when we need exactyl k and nums count.

*/
/*
  Input: nums = [1,2,1,2,3], k = 2
*/
func subarraysWithKDistinct(nums []int, k int) int {
	helper := func(k int) int {
		var mp = make(map[int]int, k)
		var l, r = 0, 0
		var count = 0
		for r < len(nums) {
			mp[nums[r]]++
			for l < len(nums) && len(mp) > k {
				mp[nums[l]]--
				if mp[nums[l]] == 0 {
					delete(mp, nums[l])
				}
				l++
			}
			count += (r - l + 1)
			r++
		}
		return count
	}
	return helper(k) - helper(k-1)
}
