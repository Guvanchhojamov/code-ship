package recursion

import "sort"

/*

Given a 0-indexed integer array nums of size n
and two integers lower and upper,
return the number of fair pairs.

A pair (i, j) is fair if:
   0 <= i < j < n, and lower <= nums[i] + nums[j] <= upper

Example 1:
Input: nums = [0,1,7,4,4,5], lower = 3, upper = 6
Output: 6
Explanation: There are 6 fair pairs: (0,3), (0,4), (0,5), (1,3), (1,4), and (1,5).

Example 2:
Input: nums = [1,7,9,2,5], lower = 11, upper = 11
Output: 1
Explanation: There is a single fair pair: (2,3).

Constraints:

1 <= nums.length <= 10^5
	nums.length == n
-10^9 <= nums[i] <= 10^9
-10^9 <= lower <= upper <= 10^9
*/
/*
  Simple brute force approach in there is
  in 2 loop generatea all pairs and check for fairity.
	tc: n*n
	sc: O(1)
   but i think this gives us TLE error. So how can we optimize?

Using hints:
1. Sort the array in ascending order.
2. For each number in the array,
   keep track of the smallest and largest numbers in the array that can form a fair pair with this number.
3. As you move to larger number, both boundaries move down.
    0,1,7,4,4,5
0 1 4 4 5 7    3,6
0 1 2 3 4 5
0-4,5
*/
/*
  We think our, sorting is mess up the indexes, but this is actually not affected for sum:
	a[i]+a[j] = a[j]+a[i] it doesn't matter where is our element we have bounds.
*/

/*
  Function - lower_bound(nums, low, high, element):
	1. Initialize a loop that continues as long as low is less than or equal to high:
		Calculate the middle index mid using the formula low + (high - low) / 2.
		If nums[mid] is greater than or equal to element, adjust the high index to mid - 1.
		Otherwise, adjust the low index to mid + 1.
		Return the low index after the loop ends, which represents the lower bound position.

Main Function - countFairPairs(nums, lower, upper):
	1.Sort the array nums.
    2.Initialize a variable ans to 0, which will hold the count of valid pairs.
    3.Iterate through each element in the sorted array using index i:
    4.For each element nums[i], determine the number of possible pairs with a sum less than lower:

 Use lower_bound to find the index of the first element in the subarray nums[i + 1] to nums[end] that is
       greater than or equal to lower - nums[i].
  Similarly, determine the number of possible pairs with a sum less than or equal to upper:
 Use lower_bound to find the index of the first element in the subarray that is greater than or equal to upper - nums[i] + 1.
The difference high - low gives the count of valid pairs with sums within the range [lower, upper] for the current element.
Update ans by adding the difference calculated.

After iterating through all elements, return the value of ans.
*/

func countFairPairs(nums []int, lower int, upper int) int64 {
	var res int64 = 0

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		l := lowerBound(nums, i+1, len(nums)-1, lower-nums[i])
		u := lowerBound(nums, i+1, len(nums)-1, upper-nums[i]+1)
		res += int64(u - l)
	}
	return res
}

func lowerBound(nums []int, low, high, target int) int {
	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
