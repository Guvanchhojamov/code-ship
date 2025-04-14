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

/*
You are given a positive integer days representing the total number of days an employee is available for work (starting from day 1).
You are also given a 2D array meetings of size n where, meetings[i] = [start_i, end_i] represents the starting and ending days of meeting i (inclusive).

Return the count of days when the employee is available for work but no meetings are scheduled.

Note: The meetings may overlap.



Example 1:

Input: days = 10, meetings = [[5,7],[1,3],[9,10]]

Output: 2

Explanation:

There is no meeting scheduled on the 4th and 8th days.

Example 2:

Input: days = 5, meetings = [[2,4],[1,3]]

Output: 1

Explanation:

There is no meeting scheduled on the 5th day.

Example 3:

Input: days = 6, meetings = [[1,6]]

Output: 0

Explanation:

Meetings are scheduled for all working days.



Constraints:

1 <= days <= 10^9
1 <= meetings.length <= 10^5
meetings[i].length == 2
1 <= meetings[i][0] <= meetings[i][1] <= days

*/
