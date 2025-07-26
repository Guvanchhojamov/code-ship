package heap

import "container/heap"

/*
215. Kth Largest Element in an Array

Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

Example 1:
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

Example 2:
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4


Constraints:

1 <= k <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
*/
/*
 We need to return k th largest element, not distinc once just kth largest in order.

 brute force:
  1. sort arr in asc or desc order start from back or front,
    iterate and descrease k until k>0; if k==0 return element and break.
 tc:nlogn+n
 sc: O(1)

2. Can we use priority queue in there?
	for example if we use maxheap.
1. build maxheap from given arr.
2. Start pop elements from heap, and decrease k.
3. if k==0 return popped element and break.
tc: n+logn = ~n
sc: O(1)
*/
/*
 ok, lets try to implement with maxheap.
  We already have implemented maxheap use in
*/

func findKthLargest(nums []int, k int) int {
	var pq = &MaxHeap{}

	heap.Init(pq)
	for _, v := range nums {
		heap.Push(pq, v)
	}
	for k > 1 {
		heap.Pop(pq)
		k--
	}

	return heap.Pop(pq).(int)
}
