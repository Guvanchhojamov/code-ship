package daily

import "sort"

/*

3487. Maximum Unique Subarray Sum After Deletion

You are given an integer array nums.

You are allowed to delete any number of elements
from nums without making it empty.
After performing the deletions, select a subarray of nums such that:

- All elements in the subarray are unique.
- The sum of the elements in the subarray is maximized.
- Return the maximum sum of such a subarray.

Example 1:
Input: nums = [1,2,3,4,5]

Output: 15

Explanation:

Select the entire array without deleting any element to obtain the maximum sum.

Example 2:

Input: nums = [1,1,0,1,1]
Output: 1
Explanation:

Delete the element nums[0] == 1, nums[1] == 1, nums[2] == 0, and nums[3] == 1.
Select the entire array [1] to obtain the maximum sum.
Example 3:
Input: nums = [1,2,-1,-2,1,0,-1]

Output: 3
Explanation:
Delete the elements nums[2] == -1 and nums[3] == -2,
and select the subarray [2, 1] from [1, 2, 1, 0, -1] to obtain the maximum sum.

Constraints:
1 <= nums.length <= 100
-100 <= nums[i] <= 100
*/
/*
 We need to take max subarray, souch that:
  - all elements are unique.
  - all elements are maximized.
  - if not unique or no needed number we can delete.
  - the result subarray cant be empty.
  - We can delete any time we want.

  Brute force:
- To take all maximized number we need to sort array in asc order.
- Start from back, comparing with next element if equal skip.
- since all max elements already in asc order they migth apper near, so
  no need check for min.
- otherwise sum elemts.
- after iterating until 0 return sum as result.
tc: nlogn + n
sc: 1. Not used any extra space.

problem with signed nums:
create distint elements map.
after sort check last element:
	if last < 0:
		for i:=0;i<n;i++:
			mp[nums[i]] = nums[i]
	else if last > 0:
		for i:=n-1;i>=0;i--:
			if nums[i]<0:
				break
			mp[nums[i]]=nums[i]
	else:
		return 0 - last element is 0 is max for all - elements.
	for v,ex:=range mp:
		res+=v
	return res
*/

func maxSum(nums []int) int {
	var res = 0
	var mp = make(map[int]int, 0)

	sort.Ints(nums)
	if nums[len(nums)-1] > 0 {

		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] < 0 {
				break
			}
			mp[nums[i]] = nums[i]
		}

	} else if nums[len(nums)-1] < 0 {
		return nums[len(nums)-1]
	} else {
		return 0
	}

	for v := range mp {
		res += v
	}
	return res
}

/*

Maximum Unique Subarray Sum After Deletion

Approach: Duplicate Removal for Positive Numbers
Intuition
The problem essentially requires finding a non-empty subsequence of non-repeating elements that gives the maximum sum.
To achieve this, we can greedily collect all positive numbers in a hash set to ensure uniqueness and then sum them.
If there are no positive numbers, we return the maximum value from the array.

Implementation

Complexity analysis
Let n be the length of the array nums.

Time complexity: O(n).

We traverse the array exactly once, and each operation on the hash table takes constant time in the average case.

Space complexity: O(n).

This comes mainly from the space used by the hash table to store the positive numbers.


*/
/*
 Optimial from editorials.
func maxSum(nums []int) int {
	positiveNumsSet := make(map[int]bool)
	maxNum := nums[0]
	for _, num := range nums {
		if num > 0 {
			positiveNumsSet[num] = true
		}
		maxNum = max(maxNum, num)
	}

	if len(positiveNumsSet) == 0 {
		return maxNum
	}
	sum := 0
	for num := range positiveNumsSet {
		sum += num
	}
	return sum
}

*/
/*
 Summary read repeatly the desc.
 1. Ask questions, cases. what if all elements are <0.
 2. What if all elements are > 0.
 3. What if 0 is the max element? same as 1.
 Then we clarify solution using this answers.
 Always do the dry run and try to find approach.

*/
