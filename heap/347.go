package heap

import (
	"container/heap"
	"fmt"
)

/*

347. Top K Frequent Elements

Given an integer array nums and an integer k,
return the k most frequent elements.
You may return the answer in any order.


Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]

Constraints:
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.

Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

*/
/*
 Given nums array. we need to return elements with frequence at most k times.
 so, freqNum >= k.
- need return all satisfing elements in array.
- tc must be better than nlogn.
- num < 0 can be.
- min len nums is =1.
*/
/*
 brute force:
1. Sort the nums.  After sort all equal elements place next to each other.
2. Iterate from front to back.
3. Count same elements to take frequency.
4. if count of freq >=k then add to the response arr and skip until next element.
tc: nlogn + n
sc: O(1) - exept result arra we dont use any extra space.
 not acceptable because tc is big. TLE.

2. Using hash map
 1. create hash map key- num. val = num frequenct count.
 2. Build map of num frequence iterating over nums.
 3. Iterate over map and if val >=k then add to response.
 4. return response array.
 tc: n+n = N better.
 sc: O(n) for hashmap.
3. Can we optimize or use other solution? Maybe priority queue?
   How we can use priority queue in this problem?

 1. Create Priority queue using value and frequency.
 2. The max frequency always in the root node.
 3. Build Priority queue from nums and frequency.
 4. Pop elements from heap and if frequency is >= k then add to response arr.

 tc: N for quque +  on each push heapify algo is logN.  nLogN. 	+ N for pop() too.

 1. solve usuign hash map approach.
 2. solve usin heap since we learn heap approach. But heap approach is slower i think.
 Need to clearify description.

 We need to return  k most frequent elements not k frequent elements.
 1. Create map, sort by frequency, and return first k nums.
 tc: n + nlogn.
 sc: N
2. Build maxHeap for key value as key-number, val=frequency.
   pop k times from heap.
   build heap - N
   pop k times - k
   tc: N+K
   sc: N for heap elements. N for map elements.
Now how to build the heap with num frequency.

0. Set heap for use freq as main value and keep heap constraint using this value.
1. Create map num:freq
2. Add to map and push to heap this element with frequncy.
3. After one iteration pop heap element and ad to result until k.
4. return result arr.
*/

type element struct {
	Num  int
	Freq int
}

type FreqHeap struct {
	Nums []element
}

func (fh FreqHeap) Len() int {
	return len(fh.Nums)
}

// sort by frequency value.
func (fh FreqHeap) Less(i, j int) bool {
	return fh.Nums[i].Freq > fh.Nums[j].Freq
}

func (fh FreqHeap) Swap(i, j int) {
	fh.Nums[i], fh.Nums[j] = fh.Nums[j], fh.Nums[i]
}

func (fh *FreqHeap) Push(x any) {
	fh.Nums = append(fh.Nums, x.(element))
}

func (fh *FreqHeap) Pop() any {
	tmp := fh.Nums
	x := tmp[len(tmp)-1]
	fh.Nums = tmp[0 : len(tmp)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	var mp = make(map[int]element)
	var res = make([]int, 0, len(nums))

	for _, v := range nums {
		if el, ok := mp[v]; ok {
			mp[v] = element{Num: el.Num, Freq: el.Freq + 1}
		} else {
			mp[v] = element{Num: v, Freq: 1}
		}
	}

	pq := &FreqHeap{}
	heap.Init(pq)
	for _, el := range mp {
		heap.Push(pq, el)
	}
	fmt.Println(k, mp, pq)
	for k > 0 {
		x := heap.Pop(pq)
		fmt.Println(x.(element).Num)
		res = append(res, x.(element).Num)
		k--
	}
	return res
}
