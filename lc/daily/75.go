package daily

/*
Given an array nums with n objects colored red, white, or blue,
sort them in-place so that objects of the same color are adjacent,
with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

Example 1:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Example 2:
Input: nums = [2,0,1]
Output: [0,1,2]


Constraints:

n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.


Follow up: Could you come up with a one-pass algorithm using only constant extra space?
*/
/*
 given nums 0-th 1-th and 2-th and, we need to sort them, in place.
 all 1-0th,2-1th, 3-2th.
 brute force:
	create 3 vars witch contains count of 0th 1th and 2th.
	- zero,ones, twos
	create empty result map and add
	while zeros > 0:
	result.add(0)
    while ones > 0:
	result.add(1)
	while twos > 0:
	result.add(2)
TC: O(n) because N/3+N/3+N/3 = N
SC: O(n) because we use result arr.
	 But we can reduce space replacing same arr nums by index.
SC: O(1) we can say.
  How can we solve optimal?
	Doing in place replacement?
*/
/*
Example 1:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Example 2:
Input: nums = [2,0,1]
Output: [0,1,2]

 We can use two pointer approach in here.
i,j,k
i=0 - 0 th keeper
j= iterator
k is 2-th keeper.
since all 0th must be on begin and all 2-th must be at the end.
i=0, k=len(nums-1)
with j we start iterate j=0 -> n-1.
  i,j,k=0,0,len(nums)-1
 for j<len(nums)-1 && j<=k:
  switch nums[j]:
case 0:
	nums[i],nums[j] = nums[j],nums[i]  // replace with last idx 0, and increment the i
	i++
case 1:
	j++
	continue
case 2:
	nums[j],nums[k] = nums[k],nums[j]
	k--
   // we keep j to check new value==0 if it is 1 in next iteration case 1 will be completed.

go trow test case:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
202110
i=0,j=0,k=5
002112
i=0, j=1,k=4
002112
i=1,j=1,k=4
002112
i=1,j=2
001122

Input: nums = [2,0,1]
Output: [0,1,2]
201
102 i=1,j=1,k=1
102
*/
func sortColors(nums []int) {
	i, j, k := 0, 0, len(nums)-1
	for j <= k {
		switch nums[j] {
		case 0:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case 1:
			j++
			continue
		case 2:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
	}
}
