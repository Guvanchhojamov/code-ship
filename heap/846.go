package heap

import "container/heap"

/*

846. Hand of Straights

Alice has some number of cards and she wants to rearrange the cards into groups
so that each group is of size groupSize, and consists of groupSize consecutive cards.

Given an integer array hand where hand[i] is the value written on the ith card and an integer groupSize,
return true if she can rearrange the cards, or false otherwise.



Example 1:

Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
Output: true
Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]
Example 2:

Input: hand = [1,2,3,4,5], groupSize = 4
Output: false
Explanation: Alice's hand can not be rearranged into groups of 4.



Constraints:

1 <= hand.length <= 10^4
0 <= hand[i] <= 10^9
1 <= groupSize <= hand.length


Note: This question is the same as 1296: https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/
*/

/*
	We have cards array, we need divide them n groups with size K.
	- size of all groups must be equal.
	- group must contain consequtive elements.
	-  What is consecutie elements?: nums witch placed in asc(increasing) order.
  - Return is Alice can divide cards exactly N groups with size K.
    true if yes, false if no.
    len(hand) = n.
    1<=n<=10^9.
    0 <= hand[i] <= 10^9.

1. Freq count and minheap.

define elements:

	size = 0
	groupsCount = 0
  - To create consecutive group, we need to start from min element.
  - start from min, to until end.
  - each time take next min element after min.
  - and increment size.
  - if size == gorupSize : inc(groupCount) and size = 0.
  - go until end.
  - after iteration:
    if len(hands) mod groupsCount==0 then return true; else return false.

Ok, how we can take min after min element efficiently? Exatcly using heap.

	use minheap.

tc: N add elemens to minheap + N * Pop min (logn); n+nlogn.
*/
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

//

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	var mp = make(map[int]int, len(hand))
	var pq = &MinHeap{}
	heap.Init(pq)
	for _, v := range hand {
		mp[v]++
	}

	for k, _ := range mp {
		heap.Push(pq, k)
	}

	for pq.Len() > 0 {
		tmp := *pq
		first := tmp[0]

		for i := first; i < first+groupSize; i++ {
			if _, ok := mp[i]; !ok {
				return false
			}
			mp[i]--
			if mp[i] == 0 {
				if i != tmp[0] {
					return false
				}
				heap.Pop(pq)
				//  delete(mp,i)
			}
			// fmt.Println(i,tmp, mp)
		}
	}

	return true

}

/*
 if len(hands) mod groupSize != 0 :
    return false.
 - count of frequency of each number. and store it in map.
 - after push all unique keys to heap.
 - iterate until pq is not empty:
    first:=pq[0]
    for i in range first+groupSize:
        if mp[i] != ok: return false.
        mp[i]--
        if mp[i] = 0:
            if i != mp[0]:  // if num with we want to pop is not min, return false
                return false.
            pq.Pop(mp[i])

    return true



*/
