package daily

/*
1290. Convert Binary Number in a Linked List to Integer

Given head which is a reference node to a singly-linked list.
The value of each node in the linked list is either 0 or 1.
The linked list holds the binary representation of a number.

Return the decimal value of the number in the linked list.
The most significant bit is at the head of the linked list.

Example 1:

Input: head = [1,0,1]
Output: 5
Explanation: (101) in base 2 = (5) in base 10

Example 2:
Input: head = [0]
Output: 0


Constraints:
- The Linked List is not empty.
- Number of nodes will not exceed 30.
- Each node's value is either 0 or 1.
*/
/*
 What we need?
- we need take number from linked list and convert it to decimal value.
- need to start from first siginficant bit, becaouse before them already 0.
bf:
- iterate over ll and insert all nums, to new arr.
- iterate over arr and convert nums to decimal.
tc: n+n
sc: N for nums array.

Can we optimize space?
 yes we don't need array, we need to iterate over linked list and take num starting
 first siginficant bit and to convert decimal operation.
 tc: O(n)
 sc: O(1)
 this is the optimal solution.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}

	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			res += 1 << (len(nums) - 1 - i)
		}
	}
	return res
}

// optimal version.
func getDecimalValue2(head *ListNode) int {
	result := 0
	for head != nil {
		result = result*2 + head.Val
		head = head.Next
	}
	return result
}
