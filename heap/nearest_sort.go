package heap

import "container/heap"

/*

Sort a nearly sorted (or K sorted) array

Given an array arr[] and a number k.

The array is sorted in a way that every element is at max k distance away from it sorted position.
It means if we completely sort the array,
then the index of the element can go from i - k to i + k where i is index in the given array.
Our task is to completely sort the array.

Examples:

Input: arr= [6, 5, 3, 2, 8, 10, 9], k = 3
Output: [2, 3, 5, 6, 8, 9, 10]

Input: arr[]= [1, 4, 5, 2, 3, 6, 7, 8, 9, 10], k = 2
Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

*/
/*
 Input: arr= [6, 5, 3, 2, 8, 10, 9], k = 3
  each element in array max k distance away from its correct position.
  it means if after swap with correct position, the element can be in. k-i or k+i position.
  6 5 3 2 8 10 9
  0 1 2 3 4 5  6
  6 - left no cant go.
  6 - go to right until, 6>a[i] or i < n and i<i+k.

  check element -1 after swap:
	if -1 > i: swap; i--
else: continue to next el.
brute force sort with buildin.
  n*logn.

using heap: use minheap to keep min element in top.
take first k elements and add to heap.

for i:=k+1 to n:
 pop; add to res; push new element a[i]
after iteration in heap remains k element;
pop them again until heap is empty. and ad to result.
return result;
tc: k+(n-k+1)logk = nlogk.
sc: O(n) for result.
*/
func sortedK(nums []int, k int) []int {
	var res = make([]int, 0, len(nums))

	pq := &MaxHeap{}
	heap.Init(pq)

	for i := 0; i <= k; i++ {
		heap.Push(pq, nums[i])
	}

	for i := k + 1; i < len(nums); i++ {
		x := heap.Pop(pq).(int)
		res = append(res, x)
		heap.Push(pq, nums[i])
	}
	for pq.Len() > 0 {
		x := heap.Pop(pq).(int)
		res = append(res, x)
	}
	return res
}

// Note: Use minheap for this problem instead of implemented maxheap
