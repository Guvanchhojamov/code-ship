package LL

/*
Given the head of a singly linked list, return the middle node of the linked list.
If there are two middle nodes, return the second middle node.


Example 1:
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.

*/
/*

 */

func middleNode(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var tmp = head
	var listLen = 0

	for tmp != nil {
		listLen++
		tmp = tmp.Next
	}

	//var med = math.Floor(float(listLen / 2))
	tmp = head
	var cnt = 0
	for tmp != nil {
		if cnt == 0 {
			break
		}
		cnt++
		tmp = tmp.Next
	}
	return tmp
}

/*
  How we can see this is the brute force solution, we need optimize it using slow and fast pointers.
   slow pointer iterates jumping 1 node.
   fast pointer iterates jumping 2 nodes.
   when fast pointer reached, the enf of the list, the slow pointer stops in middle of the list.
    Because slow is 2x slower than fast so, fast /2 = slow.
   fast = 0,2,4,5
   slow = 0,1,2,3
    0 1 2 3 4 5
     we have 2 middles and 3 is the
   what's fasts end point?
     fast == nil
   0 1 2 3 4
  f: 0, 2 , 4
  s: 0 1 2

*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead = reverseList(head.Next)
	var nxt = head.Next
	nxt.Next = head
	head.Next = nil

	return newHead
}

//  1 2 3 4 5 -> nil
// nil
