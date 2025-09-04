package heap

import "container/heap"

/*
295. Find Median from Data Stream
The median is the middle value in an ordered integer list.
If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

For example, for arr = [2,3,4], the median is 3.
For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
Implement the MedianFinder class:

MedianFinder() initializes the MedianFinder object.
void addNum(int num) adds the integer num from the data stream to the data structure.
double findMedian() returns the median of all elements so far.
Answers within 10^-5 of the actual answer will be accepted.


Example 1:

Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]

Explanation
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0


Constraints:

-10^5 <= num <= 10^5
There will be at least one element in the data structure before calling findMedian.
At most 5 * 104 calls will be made to addNum and findMedian.


Follow up:

If all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?
If 99% of all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?
*/
/*
 What we need ?
	We have big array.
	We need to find median number from this array. If length is,
	 if odd: median is n/2 index number.
	 if even:  (n/2 + (n/2)-1) / 2.

 How do we find fast median number of array ?
 [3,2,4,5,6] -
  0 1 2 3 4  - 5/2-1 = 2 => 4.

  3,2,5,6,7,8,
  0 1 2 3 4 5 - 6/2 + 5/2 = 3+2 => 5+6 = 11/2 = ~2.5

1. approach.
	- init ds with nums []int array to store stream nums.
	- addNum, adds num to ds.nums array.
	- find median:
		if len(nums) is odd:
			return len(num)/2. sa float number.
		if len is even:
			return (n/2+(n/2-1)) /2 as float type.
	tc: O(1)
	sc: O(1)
	Why this problem is hard? idont know.
Why this is hard well, its not the simple.
if read carefully the array should be sorted always, in asc order.
so on each addition we need to sort.  or use efficient data structure.
brute force:
	sort each time array.
	 [3,2,4,5,6]
	 2,3,4,5,6

	 6,5
	 2,4,3

	 if we separate array 2 parts?
	 1. maxheap.
	 2. minheap.
	 len of maxheap is always: n/2
	 len of minheap is always: n/2:n
	 if len of whole array is odd:
	 	take top of maxheap.
	if len is even:
		take root from maxheap + minheap /2.
	tc: n for build max, minheap..
	sc: n form max+min heap.
*/

type MedianFinder struct {
	Largeq MinHeap
	Smallq MaxHeap
}

func Constructor3() MedianFinder {
	largeq := &MinHeap{}
	smallq := &MaxHeap{}
	heap.Init(largeq)
	heap.Init(smallq)
	return MedianFinder{
		Largeq: *largeq,
		Smallq: *smallq,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// Push into Smallq (max-heap by default)
	heap.Push(&this.Smallq, num)

	// Smallq top <= Largeq top
	if this.Smallq.Len() > 0 && this.Largeq.Len() > 0 && this.Smallq[0] > this.Largeq[0] {
		val := heap.Pop(&this.Smallq).(int)
		heap.Push(&this.Largeq, val)
	}

	if this.Smallq.Len() > this.Largeq.Len()+1 {
		val := heap.Pop(&this.Smallq).(int)
		heap.Push(&this.Largeq, val)
	} else if this.Largeq.Len() > this.Smallq.Len()+1 {
		val := heap.Pop(&this.Largeq).(int)
		heap.Push(&this.Smallq, val)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	smalls := this.Smallq.Len()
	larges := this.Largeq.Len()

	if smalls > larges {
		return float64(this.Smallq[0])
	} else if larges > smalls {
		return float64(this.Largeq[0])
	}
	return (float64(this.Smallq[0]) + float64(this.Largeq[0])) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
