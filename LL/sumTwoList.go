package LL

/*
You are given two non-empty linked lists representing two non-negative integers.

The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.

*/
/*
  1. Make it like simple sum how we can do in the paper.
      Keep track carry as a variable.
    and sum each number from iterating head of each list.

	for (l1 != nil and l2 != nil):
    	nodeSum = l1+l2+carry.
		carry = nodeSum / 10
		nodeVal = nodeSum % 10
       newNode.Val = nodeVal
       l1.Next,
       l2.Next
	   newNode.Next

	for l1!=nil and carry>0:
		nodeSum = l1+carry.
		carry = nodeSum / 10
		nodeVal = nodeSum % 10
		l1.Next,
	   	newNode.Next
  return newNode.Next

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	var current = dummy
	var tmp1 = l1
	var tmp2 = l2
	var carry = 0

	for tmp1 != nil || tmp2 != nil || carry != 0 {
		val1, val2 := 0, 0

		if tmp1 != nil {
			val1 = tmp1.Val
			tmp1 = tmp1.Next
		}
		if tmp2 != nil {
			val2 = tmp2.Val
			tmp2 = tmp2.Next
		}

		nodeSum := val1 + val2 + carry
		carry = nodeSum / 10
		nodeVal := nodeSum % 10

		current.Next = &ListNode{Val: nodeVal}
		current = current.Next
	}
	return dummy.Next
}
