package LL

/*
Given the head of a linked list, rotate the list to the right by k places.

Example 1:
Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]

Example 2:
Input: head = [0,1,2], k = 4
Output: [2,0,1]


Constraints:

The number of nodes in the list is in the range [0, 500].
-100 <= Node.val <= 100
0 <= k <= 2 * 10^9
*/

/*
   Brute force:
- convert to arr
- rotate array, right for k place.
- create new list
  tc: n+n+n = 3n
  sc: n+n = 2n
 Constraints are very big we got tle.
2. For each k find last element, link it to head, and prev link to nil.
   k*n = n*2
 Optimal approach?
    make k = k %n to reduce cycle rotates.
   Take first k nodes to temp.
   take tail of the node.
   link tail to temp nodes head.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	temp := head
	length := 1
	for temp.Next != nil {
		length++
		temp = temp.Next
	}

	temp.Next = head

	k = k % length
	end := length - k

	for end > 0 {
		temp = temp.Next
		end--
	}

	newHead := temp.Next
	temp.Next = nil

	return newHead
}

/*
  get list length
  create temp list with k nodes

*/
