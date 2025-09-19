package sliding_window

/*

904. Fruit Into Baskets
You are visiting a farm that has a single row of fruit trees arranged from left to right.
The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.

You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
Given the integer array fruits, return the maximum number of fruits you can pick.



Example 1:

Input: fruits = [1,2,1]
Output: 3
Explanation: We can pick from all 3 trees.
Example 2:

Input: fruits = [0,1,2,2]
Output: 3
Explanation: We can pick from trees [1,2,2].
If we had started at the first tree, we would only pick from trees [0,1].
Example 3:

Input: fruits = [1,2,3,2,2]
Output: 4
Explanation: We can pick from trees [2,3,2,2].
If we had started at the first tree, we would only pick from trees [1,2].


Constraints:

1 <= fruits.length <= 10^5
0 <= fruits[i] < fruits.length

*/

/*
 We need to fill buckets with fruits as much as it can be possible max.
 - have only 2 backet.
 - when the fuit type cant in bucket we must stop.
 - return the max.

 brute force:
 - generate all combinations anc check with maximum each time.

	Input: fruits = [1,2,3,2,2]
	mp = map[int]int. key- number; val - count.
	i=0->n:
	j=i->n:
	 if len(mp) > 2: break
	 mp[num]++
	 mx = max(mx, mp[0]+mp[1])
	 return mx
	tc: n*n
	sc: O(1) - because the map lenth is constant and always==2.
	We got TLE error
*/
/*
 How we can optimze?
 - use 2 potiner to keep window.
 - keep counter.
 - when len.map > k: increrase left. remove left from map. decrease counter.
 - else add to map right increase right, increase counter.
 - check mx,counter.
 - go until right > n;
 - return mx.
*/

func totalFruit(fruits []int) int {
	var mp = make(map[int]bool, 2)
	var count = 0
	var left, right = 0, 0
	var mx = 0
	for right < len(fruits) {
		mp[fruits[right]] = true
		if len(mp) > 2 {
			delete(mp, fruits[left])
			left++
			count--
		}
		count++
		mx = max(mx, count)
	}
	return mx
}
