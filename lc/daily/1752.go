package daily

/*
1752. Check if Array Is Sorted and Rotated
Given an array nums, return true if the array was originally sorted in non-decreasing order, then rotated some number of positions (including zero). Otherwise, return false.

There may be duplicates in the original array.

Note: An array A rotated by x positions results in an array B of the same length such that A[i] == B[(i+x) % A.length], where % is the modulo operation.

Example 1:

Input: nums = [3,4,5,1,2]
Output: true
Explanation: [1,2,3,4,5] is the original sorted array.
You can rotate the array by x = 3 positions to begin on the the element of value 3: [3,4,5,1,2].
Example 2:

Input: nums = [2,1,3,4]
Output: false
Explanation: There is no sorted array once rotated that can make nums.
Example 3:

Input: nums = [1,2,3]
Output: true
Explanation: [1,2,3] is the original sorted array.
You can rotate the array by x = 0 positions (i.e. no rotation) to make nums.

Constraints:

1 <= nums.length <= 100
1 <= nums[i] <= 100
*/

// func main() {

// }

func check(nums []int) bool {
	var breakPoint = 0
	var n = len(nums) - 1 // it is const value
	for i := 0; i < n; i++ {
		if nums[i] > nums[i+1] {
			breakPoint++
		}
	}
	if nums[0] < nums[n] {
		breakPoint++
	}

	return breakPoint <= 1
}

/*
   We need to know it that, is array sorted and rotated.
   key points:
      if array sorted and rotated,
      - a[0]>=a[n] always, For rotation != n
      - array has always 0 or 1  breakpoints.
      breakpoint = a[i]> a[i+1], because array is non decreasing soted.
*/
