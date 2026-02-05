package dp

import (
	"math"
	"sort"
)

/*
Given a wooden stick of length n units.
The stick is labelled from 0 to n.
For example, a stick of length 6 is labelled as follows:

Given an integer array cuts where cuts[i] denotes a position you should perform a cut at.

You should perform the cuts in order, you can change the order of the cuts as you wish.

The cost of one cut is the length of the stick to be cut, the total cost is the sum of costs of all cuts.
- When you cut a stick, it will be split into two smaller sticks
(i.e. the sum of their lengths is the length of the stick before the cut). Please refer to the first example for a better explanation.

Return the minimum total cost of the cuts.


/*
Ok, we given, stick of len n,
And we are given arr[] where given points to cut.
We need to cut stick from this points. Each cut have its own cost.
We need to return minimum possible cost.

Input:
	arr []int, n int
Output:
	minCost int

Constraints:
arr  len: max<n; min=0;
arr[i]: max<n min:0;
Ok do arr is always valid? Can we cut always? Yes.
Is the cuts array are distinct? Yes we dont have duplicated cut points.

What is the cost?
 The cost is len of stick to be cut.
For example:
[1,2,3] - the [2] we cut from 2; the cost =3.

Let's take some examples:
 n=6;  arr = [3,1,2,5]
0,1,2,3,4,5,6;

cost: 6 + 3+2 + 4 = 15;
0,1,2 -> 0 | 1,2 -> 1 | 2;
3,4,5,6 -> 3,4 | 5,6
Ok if we start from 1?
Cost: 6+3+2+1 = 12 - for example similar like this.

*/
/*
 To do this we need to find optimal order for cutting, reordering the cuts array.
How we can do this?
 We can try all possible ways starting from i=0; and end.
So, how we do this?
For each starting point we have len(cut)-2 possible ways. Without number of itself.
And we have len(cuts) possible starting points.
So we need to try all possible ways,
How we can try all possible ways?
To try all possible ways we need to use recursion.
But how we can use recursion?
Base case
Recursion case.
What parameters we need to do this?
 n=6;  cuts = [3,1,2,5]
0,1,2,3,4,5,6;
[3,1,2,5]
 0 1 2 3
 i
we need starting point always.
idx - cut index
To dont be conflict for example, we cut from 3, and after left partition we wna t try cut from 5. We cannot able to do this because we dont have 5 len stick, so first we must to sort the cuts array in increasing order to make sure, the left part on the left part and right part on the right part always.
 [1,2 | 3,5]
  i       j
and we need to define left and right points, and take between them
We try for each cut:
Calculate cost for this cut
Cost = j-i
Left: get left cost calling recursion:
Try cut every position beetween i and j.

*/
/*
[1,2, 3,5]
 0,1,2,3,4,5,6
*/

func minCost(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)

	l, r := 1, len(cuts)-2
	m := len(cuts)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return cutter(l, r, cuts, dp)
}

func cutter(l, r int, cuts []int, dp [][]int) int {
	if l > r {
		return 0
	}
	if dp[l][r] != -1 {
		return dp[l][r]
	}
	minCost := math.MaxInt
	currCost := cuts[r+1] - cuts[l-1]
	for idx := l; idx <= r; idx++ {
		cost := currCost + cutter(l, idx-1, cuts, dp) + cutter(idx+1, r, cuts, dp)
		minCost = min(minCost, cost)
		dp[l][r] = minCost
	}
	return dp[l][r]
}

/*
0,1,2,3
[1,3]
0,1
2,3

0,1
0
1
*/
