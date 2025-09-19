package heap

import "container/heap"

/*
239. Sliding Window Maximum
You are given an array of integers nums,
there is a sliding window of size k which is moving from the very left of the array to the very right.
You can only see the k numbers in the window.
Each time the sliding window moves right by one position.

Return the max sliding window.



Example 1:

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Example 2:

Input: nums = [1], k = 1
Output: [1]


Constraints:

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
*/
/*
 brute force, check each sliding window element and return max.
 for each element, it can check and find max.
 tc: n*k.
 sc: 1
  Other we can use sliding window and store max value and check

Can we use max heap in there to keep track max element of window?
Yes we can.
1. Build max heap for k window elements.
2. if heap size is < k then get top element and push new element to window.
3. If heap size is > k then pop top element
4. add poped elements to response arr.

tc:
 build window heap - K
 iterate over arr - N
 push or pop element = logN
 k+N+2logn
sc:
 k to store window elements in heap.

*/

type Element struct {
	Value int
	Index int
}

type WindowHeap []Element

func (h WindowHeap) Len() int           { return len(h) }
func (h WindowHeap) Less(i, j int) bool { return h[i].Value > h[j].Value }
func (h WindowHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *WindowHeap) Push(x any)        { *h = append(*h, x.(Element)) }
func (h *WindowHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (p *WindowHeap) Top() any {
	tmp := *p
	if len(tmp) > 0 {
		return tmp[0]
	}
	return 0
}

func maxSlidingWindow(nums []int, k int) []int {
	var res = make([]int, 0, len(nums)/2)
	var pq = &WindowHeap{}

	heap.Init(pq)
	for i := 0; i <= k; i++ {
		heap.Push(pq, Element{Value: nums[i], Index: i})
	}
	arr := *pq
	res = append(res, arr[0].Value)
	for i := k; i < len(nums); i++ {
		heap.Push(pq, Element{Value: nums[i], Index: i})
		arr := *pq
		for arr[0].Index <= i-k {
			heap.Pop(pq)
		}
		res = append(res, pq.Top().(Element).Value)
	}

	return res
}
