package heap

import "container/heap"

/*
1331. Rank Transform of an Array.

Given an array of integers arr, replace each element with its rank.

The rank represents how large the element is. The rank has the following rules:

Rank is an integer starting from 1.
The larger the element, the larger the rank. If two elements are equal, their rank must be the same.
Rank should be as small as possible.


Example 1:

Input: arr = [40,10,20,30]
Output: [4,1,2,3]
Explanation: 40 is the largest element. 10 is the smallest. 20 is the second smallest. 30 is the third smallest.

Example 2:
Input: arr = [100,100,100]
Output: [1,1,1]
Explanation: Same elements share the same rank.

Example 3:
Input: arr = [37,12,28,9,100,56,80,5,12]
Output: [5,3,4,2,8,6,7,1,3]


Constraints:

0 <= arr.length <= 10^5
-10^9 <= arr[i] <= 10^9

*/
/*
 we need to get rank for each element.
 - max element has bigger rank,
 - rank starts from 1, is most big rank.
 - if elements are same, ranks also must be same.
 - elements can be >0 and <0
 brute force:
	- sort the arr in asc order.
	- craete map and fill map with ar relements:
		key - number, val - index.
	- iterate over arr, and take num index from map for num.
	- add to response arr. return response arr.
 tc: nlogn+n+n.
 sc: n.

 How can we optimze?
	- build max heap from given arr.
	- crate map to store number and rank.
	- start from rank=1.
	- pop from pq,
		 if mp[num] not in map set rank and increment rank.
		 add num with rank to map.
		 if not in map  set num = mp[num] rank not increment rank.
	- iterate over arr again, and call eah num from map[num].
	- add to response return response arr.
	tc: n+n+n = 3n.
	sc: n +n  = n.
*/

func arrayRankTransform(arr []int) []int {
	var res = make([]int, 0, len(arr))
	var mp = make(map[int]int)
	var pq = &MaxHeap{}
	heap.Init(pq)

	for _, v := range arr {
		heap.Push(pq, v)
	}
	rank := 0
	for pq.Len() > 0 {
		x := heap.Pop(pq).(int)
		if r, ok := mp[x]; ok {
			rank = r
		} else {
			rank++
		}
		mp[x] = rank
	}
	for _, v := range arr {
		res = append(res, mp[v])
	}
	return res
}
