package LL

/*
142. Linked List Cycle II
Given the head of a linked list, return the node where the cycle begins.
If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed).
It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.

Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.

Example 2:
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.

Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.

Constraints:

The number of the nodes in the list is in the range [0, 104].
-10^5 <= Node.val <= 10^5
pos is -1 or a valid index in the linked-list.


Follow up: Can you solve it using O(1) (i.e. constant) memory?

*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var nodesMap = make(map[*ListNode]bool, 0)
	for head != nil {
		if _, ok := nodesMap[head]; ok {
			return head
		}
		nodesMap[head] = true
		head = head.Next
	}
	return nil
}

/*
  head = [3,2,0,-4], pos = 1
  3,2,0,-4, '2'
*/
/*
   How we can optimize using 1 sc solution?
    1 2 3 4 5 6   pos = 2
    1,1
    2,3
    3,5
    4,2
    5,4
    6,6
    when fast == slow {returnfast.Next }
*/

func detectCycleII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var slow = head
	var fast = head
	var tmpHead = head
	for fast != nil && fast.Next != nil {
		if slow == fast {
			slow = tmpHead
			for slow == fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}

/*
 Given the head of a linked list,
 determine whether the list contains a loop.
 If a loop is present, return the number of nodes in the loop, otherwise return 0.
*/

/*
 Brute force.
    add node to hash map.
	node:index
	if node in map:
		tmp = node
		res = tmp.Index+1 - node.index
		node = node.Next
	return res.

    0 1 2 3 4 5  = 6
    1 2 3 4 5 6
2. Optimal.
	Find start of loop using slow and fast.
		slow == fast
	Then init
	count=1, fast = fast.Next
	while fast != slow {
		fast = fast.Next
		cnt++
	}
return cnt
*/
