package main

func main() {

}

/*
  You are given an integer array nums,
  an integer k, and an integer multiplier.

You need to perform k operations on nums.
In each operation:

 - Find the minimum value x in nums.
   If there are multiple occurrences of the minimum value,
   select the one that appears first.
 - Replace the selected minimum value x with x * multiplier.
Return an integer array denoting the final state of nums after performing all k operations.

Example 1:

Input: nums = [2,1,3,5,6], k = 5, multiplier = 2

Output: [8,4,6,5,6]
After operation 1	[2, 2, 3, 5, 6]
After operation 2	[4, 2, 3, 5, 6]
After operation 3	[4, 4, 3, 5, 6]
After operation 4	[4, 4, 6, 5, 6]
After operation 5	[8, 4, 6, 5, 6]
*/

/*
  1. The first solution coming to mind is.
     create helper function witch,
	finds min value from array, and mutiple with multiplier and returns the changed array.
	Call this function recursively,
	while count<=k.
	if count==k return array.

*/

func getFinalState(nums []int, k int, multiplier int) []int {
	counter := 1
	return findMaxAndmutiple(nums, counter, k, multiplier)
}

func findMaxAndmutiple(nums []int, counter, k, multiplier int) []int {
	if counter == k {
		return nums
	}
	minNum := nums[0]
	minI := 0
	for i := range nums {
		if nums[i] < minNum {
			minNum = nums[i]
			minI = i
		}
	}
	nums[minI] *= multiplier
	return findMaxAndmutiple(nums, counter+1, k, multiplier)
}

/*

Input: nums = [2,1,3,5,6], k = 5, multiplier = 2
1:[2,2,3,5,6]
2:[]
*/
