package daily

/*
You are given a binary array nums.
You can do the following operation on the array any number of times (possibly zero):
 - Choose any 3 consecutive elements from the array and flip all of them.
Flipping an element means changing its value from 0 to 1, and from 1 to 0.

Return the minimum number of operations required to make all elements in nums equal to 1.
If it is impossible, return -1.



Example 1:
Input: nums = [0,1,1,1,0,0]
Output: 3

Explanation:
We can do the following operations:
Choose the elements at indices 0, 1 and 2. The resulting array is nums = [1,0,0,1,0,0].
Choose the elements at indices 1, 2 and 3. The resulting array is nums = [1,1,1,0,0,0].
Choose the elements at indices 3, 4 and 5. The resulting array is nums = [1,1,1,1,1,1].

Example 2:
Input: nums = [0,1,1,1]
Output: -1
Explanation:
It is impossible to make all elements equal to 1.

Constraints:
3 <= nums.length <= 10^5
0 <= nums[i] <= 1
*/

/*
   n = [0,1,1,1,0,0]
        1 0 0 1 0 0
		1 1 1 0 0 0
		1 1 0 1 1 0
		1 1 0 1 1 0
        1
 1. brute force:
   1 1 0 0 0 1 1 0 1 0 0 1.   12
       1 1 1
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func minOperations(nums []int) int {
	var cnt = 0
	var i = 0
	for i < len(nums)-2 {
		if nums[i] == 0 {
			nums[i] = 1
			nums[i+1] = swapBit(nums[i+1])
			nums[i+2] = swapBit(nums[i+2])
			cnt++
		}
		i++
	}
	for _, v := range nums[len(nums)-2:] {
		if v == 0 {
			return -1
		}
	}
	return cnt
}

func swapBit(b int) int {
	if b == 0 {
		return 1
	}
	return 0
}
