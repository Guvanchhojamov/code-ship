package heap

import "container/heap"

/*
703. Kth Largest Element in a Stream
You are part of a university admissions office and need to keep track of the kth highest
test score from applicants in real-time.
This helps to determine cut-off marks for interviews and admissions dynamically
as new applicants submit their scores.

You are tasked to implement a class which,
for a given integer k,
maintains a stream of test scores and continuously returns the kth highest test score after
a new score has been submitted.
More specifically, we are looking for the kth highest score in the sorted list of all scores.

Implement the KthLargest class:

KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of test scores nums.
int add(int val) Adds a new test score val to the stream and returns
the element representing the kth largest element in the pool of test scores so far.


Example 1:

Input:
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

Output: [null, 4, 5, 5, 8, 8]
8 5 4 3 2 - 4
8 5 5 .. -5
10 8 5 5 ... - 5
10 9 8 ... - 8
10 9 8 ... - 8

Explanation:

KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3); // return 4
kthLargest.add(5); // return 5
kthLargest.add(10); // return 5
kthLargest.add(9); // return 8
kthLargest.add(4); // return 8

Example 2:

Input:
["KthLargest", "add", "add", "add", "add"]
[[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]

Output: [null, 7, 7, 7, 8]

Explanation:

KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);
kthLargest.add(2); // return 7
kthLargest.add(10); // return 7
kthLargest.add(9); // return 7
kthLargest.add(9); // return 8


Constraints:

0 <= nums.length <= 10^4
1 <= k <= nums.length + 1
-10^4 <= nums[i] <= 10^4
-10^4 <= val <= 10^4
At most 10^4 calls will be made to add.
*/

/*
 We neew to create stream of nums scores .
	- init
	- add new score.
   - each add need return k'th greater score as result.
 - the range is big so need to solve this with optimal solution.
  What we can do?
	1. Create Scores class with nums array(scores).
		sort and init array.
	2. On each add we need:
		- add val to nums arr.
		- sort nums in desreasing order.
		- return k'th (or k-1 if 0 indexed) element from front.
	tc:  nlong+ n * nlogn = ~ n^3.
	sc: 1. We dont use any extra space.
But with this solution we got TLE.

2. How can we optimize?
  We can avoid sorting each time if we effiently store scores array keeping order.
	How? We can store scores in MaxHeap- with store bigger element in the root and
		heapify other elements on each push or pop.
	ok lets write step-by step approach.
- init nums array. Heapify nums array.
- add:
	- push new, automatically heapifyed.
	- copy arr.
	- pop until len>0 and k>0. Decrease k each pop.
- return k'th popped element as result of add.
	*init - iterate over nums and create heap. - N.
	*add - push - logn + k*logn. k times pop.
	tc: n+logn+kLogn.
	sc: N for maxheap. and copy.



*/

type KthLargest struct {
	K    int
	Nums MinHeap
}

func Constructor2(k int, nums []int) KthLargest {
	var pq = &MinHeap{}
	heap.Init(pq)
	for i := range nums {
		heap.Push(pq, nums[i])
	}
	return KthLargest{
		K:    k,
		Nums: *pq,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.Nums, val) // add new element.
	for this.Nums.Len() > this.K {
		heap.Pop(&this.Nums)
	}

	return this.Nums[0]
}

// 4, 5, 8, 2  k =3
/*
	8 5 4 2
	+10
	10 8 5 2 4
5 8 10in above approach we got time limet err.
To optimize we use heap efficiently, use minheap instead of maxheap.
	- use minheap instead max.
	- keep heap size always k.
	- i size exids k remove root element.
	return root element as kTh greatest.
Why this is works? why min heap?
   if we always keep k elements the root min element is kth largest element.
   because, the max in the last then greater+greater elements.
   all elements after root are greater from him.
   We dont neet order we need just kth greater element.

*/
