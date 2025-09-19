package heap

import (
	"container/heap"
)

/*
23. Merge k Sorted Lists

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.


Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted linked list:
1->1->2->3->4->4->5->6

Example 2:
Input: lists = []
Output: []

Example 3:
Input: lists = [[]]
Output: []


Constraints:

k == lists.length
0 <= k <= 104
0 <= lists[i].length <= 500
-104 <= lists[i][j] <= 104
lists[i] is sorted in ascending order.
The sum of lists[i].length will not exceed 104.
*/
/*
 We need to merge sorted linked lists into one sorted linked lisit.
 - It may K lists.
 - all lists already sorted seperately.
 - Need to merge them keeping sorted order.

brute force:
	- iterate over all k lists. k
	- add to array all elements of linked list.
	- sort array. nlongn
	- iterate over array and create other linked list. n
	- return the head of this linked list.
	tc:k+nlogn+n = 2n+nlogn = n +nlogn.
	sc: n+n = 2n = n
*/
/*
1.  merge like 2 soretd list, until all lists merged.
   result & ith list. result starts with empty [].
2. using dummy head and each time take min fdata from all heads,
	- create dummy head.
	- take min head from all,
	- link dummy to this, anc i-thhead = ithhead.next.
		continue again.
3. use min heap to take efficiently min node. Use nodes minheap.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type NodePQ []*ListNode

func (h NodePQ) Len() int {
	return len(h)
}

func (h NodePQ) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h NodePQ) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodePQ) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodePQ) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {

	var pq = &NodePQ{}

	heap.Init(pq)

	var dummyHead = &ListNode{}
	for _, list := range lists {
		if list != nil {
			heap.Push(pq, any(list))
		}
	}

	curr := dummyHead
	for pq.Len() > 0 {
		node := heap.Pop(pq).(*ListNode)

		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
		curr = node
		curr = curr.Next
	}
	return dummyHead.Next
}
