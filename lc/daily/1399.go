package daily

import "math"

/*
1399. Count Largest Group

You are given an integer n.
Each number from 1 to n is grouped according to the sum of its digits.

Return the number of groups that have the largest size.

Example 1:
Input: n = 13
Output: 4
Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
[1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9].
There are 4 groups with largest size.

Example 2:
Input: n = 2
Output: 2
Explanation: There are 2 groups [1], [2] of size 1.

Constraints:
1 <= n <= 10^4
*/

/*
 1. simple brute force approach?
    we need group the same elements, and count them

from 1 to n

	how we can do it using hash map ?
	key - num  (sum of the digits)
	value - frequency.

after fill the map return max elements count.

	in one loop.
	tc: n+n = 2n=n
	sc: n

Can we optimize this?
1. create another arr.
2. fill this array with sum of digits.
3. sort the arr.
4. count from back until element are not same, because the max values in back of sorted arr.
tc: n + nlogn +n
sc: n

Can we do binary search for sorted arr?
yes, but it does not much affect for tc:
lets assume we do uppper bound for sorted arr.
tc: n+nlogn+logn
sc: n
the same result so, lets use the first solution.

	So, until n<10 the sum of element digits size are 1, so
	 if n<10 we can return n
*/
func countLargestGroup(n int) int {
	if n <= 10 {
		return n
	}
	var sumsMap = make(map[int]int, 0)

	for i := 0; i <= n; i++ {
		sum := sumDigits(i)
		sumsMap[sum]++
	}
	var mx = math.MinInt
	var res = 0
	for _, val := range sumsMap {
		if val > mx {
			mx = val
			res = 1
		} else if val == mx {
			res++
		}
	}
	return res
}

func sumDigits(n int) int {
	var s = 0
	for n > 0 {
		s = s + n%10
		n = n / 10
	}
	return s
}
