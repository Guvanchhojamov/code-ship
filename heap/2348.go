package heap

/*
 2348. Number of Zero-Filled Subarrays
Given an integer array nums,
return the number of subarrays filled with 0.

A subarray is a contiguous non-empty sequence of elements within an array.


Example 1:

Input: nums = [1,3,0,0,2,0,0,4]
Output: 6
Explanation:
There are 4 occurrences of [0] as a subarray.
There are 2 occurrences of [0,0] as a subarray.
There is no occurrence of a subarray with a size more than 2 filled with 0. Therefore, we return 6.

Example 2:
Input: nums = [0,0,0,2,0,0]
Output: 9
Explanation:
There are 5 occurrences of [0] as a subarray.
There are 3 occurrences of [0,0] as a subarray.
There is 1 occurrence of [0,0,0] as a subarray.
There is no occurrence of a subarray with a size more than 3 filled with 0. Therefore, we return 9.

Example 3:
Input: nums = [2,10,2019]
Output: 0
Explanation: There is no subarray filled with 0. Therefore, we return 0.


Constraints:

1 <= nums.length <= 10^5
-109 <= nums[i] <= 10^9

*/
/*
 We are given nums array.
	- need to find subarrays, witch all elements are 0 of fileed with 0.
	- return count of the filled subarrays with 0.
[0,0,0,2,0,0]  res = 9.
 0 1 2 3 4 5
1. Gerate all subarrays and count if filled with 0.
   in two loops.
	for i=0 to n:
		for j=i to n:
			if a[j]== 0:
				res+=(j-i+1)
				continue
			else if j>i:
				i=j
			break.


for each 0 value need to calculate result += this end index assume this is the end index of subarray.
 0 0 0
 0 1 2
 1+2+3 = 6 + 1+2 = 9
 return 9

  0 0
  0 1
  1+2 = 3.
  on each occurence of 0 we need to calculate the count of subarrays until this element like this:
	endI-startI+1 and add it to result.
	for example from index 0 -> 2 we have 6 subarrays witch we may generate.
  tc: n*n.
  sc: 1

2. How w can optimize the  time complexity?
	for example using pointers?
	i,j=0,0.
	while i < n && j < n:
		while a[j] == 0 and j < n:
			res+=j-i+1
			j++
		while a[i] == 0 and i<n:
			i++
		i++, j++

	return result.
[0,0,0,2,0,0]
tc: n+n n for i n for j.
sc: 1.

*/

func zeroFilledSubarray(nums []int) int64 {
	var i, j = 0, 0
	var res = 0
	for i < len(nums) && j < len(nums) {
		for j < len(nums) && nums[j] == 0 {
			res += (j - i + 1)
			j++
		}
		for i < len(nums) && nums[i] == 0 {
			i++
		}
		i++
		j++
	}
	return int64(res)
}
