package daily

/*

904. Fruit Into Baskets

You are visiting a farm that has a single row of fruit trees arranged from left to right.
The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.

You want to collect as much fruit as possible.
 However, the owner has some strict rules that you must follow:

- You only have two baskets, and each basket can only hold a single type of fruit.
 There is no limit on the amount of fruit each basket can hold.
- Starting from any tree of your choice,
 you must pick exactly one fruit from every tree (including the start tree) while moving to the right.
 The picked fruits must fit in one of your baskets.

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

*/

/*
 we have only 2 baskets.
 types of fruits in each busket must be same.
 when we reach to the end or diff type from baskets, must stop.
 Pick max possible trees.
 fruits length are big.

To get max tree substr we need to find max leng subscrting. with combination of only 2 nums.
1. Brute force.
	1. start from each tree pick elements, and create map.
	2. compare with max after each iteration.
	return max.
tc: n*n
sc: n  We got TLE

We cant sort the arr.
- When we need to keep track some cons. in array and keep max possible value.
- We can use two pointers to keep left and right bundaries,
- we can use sliding window to move and update window elements.

fruits = [1,2,3,2,2]
Output: 4
[0,1,2,2]

ok pick i,j left and right pointer to 0,0 starting position.
  take hash map with len=2 to keep buckets.  key=number; val=number freq.
  start moving j and add to map = nums[j]
  if len(map)>2 then remove element from map nums[i]. and i++
	sum--
 else:
	sum++
	res = max(res, sum)
 until end
 tc: n
 sc: 1 for hashmap.

 [3,3,3,1,2,1,1,2,3,3,4]
3:3
2:1
*/

func totalFruit(fruits []int) int {
	var mp = make(map[int]int)
	var count = 0
	var res = 0
	var i, j = 0, 0

	for j < len(fruits) {
		mp[fruits[j]]++
		for len(mp) > 2 {
			mp[fruits[i]]--
			if mp[fruits[i]] == 0 {
				delete(mp, fruits[i])
			}
			i++
			count--
		}
		count++
		res = max(res, count)
		j++
	}
	return res
}
