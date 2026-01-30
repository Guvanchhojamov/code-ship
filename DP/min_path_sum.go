package dp

/*
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
Note: You can only move either down or right at any point in time.


Ok we have matrix with positive nums, need to find minimized sum path from 0,0 -> n-1, m-1.
Input:
Grid [][]int.
N,m int
Output:
Minimized sum of path.

Constraints:
All nums > 0.
Min len of grid is 1<=grid.leng<=n.
Can move only down, right.

Ok, Can we use graph with 2 directions, maybe we can.
And find shortest path from 0,0 to n,m.
It seems we can but, lets check apraocj and TC.
Approach:
	For this we use BFS with PriorityQueue for shortest path.
Implement pq by node value.
Push start node to pq, its 0,0, value.
Pop from pq min value cell,
And push right, bottom child.
Keep and update visited array.
Keep and update destinations array.
If popped is n-1,m-1 then break loop.
Return desctinations n-1, m-1 value.
!! we need sum destination value also.
TC: O(n*mLog(n*m)) we pop and push N*M times in worth case.
SC: O(N*M)in worth case we store n*m we check all nums in PQ.
Is this optimal? Maybe.

Lets think in other way. Maybe we can use DP there?
What if we start from destination?
We can use top-down recursion starting from dest.->start.
And choosing the min val on each move.
For recursion we need 2 cases:
// base case
 if i==0 and j==0 {return grid[i][j]}
 if i<0 or j<0 return (1e9)

//recursion case.
 up:=fn(i-1,j)
 left:=fn(i,j-1)
return grid[i][j] + min(up,left)
TC: O(N*M)
SC: O(N*M).
Ok, lets try implement this.
Any edge case?
*/

func minPathSum(grid [][]int) int {
	//edge cases
	if len(grid) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			up, left := 0, 0
			if i > 0 {
				up += grid[i-1][j]
			}
			if j > 0 {
				left += grid[i][j-1]
			}
			dp[i][j] = grid[i][j] + min(up, left)
		}
	}

	return dp[n-1][m-1]
}

/*
 Ok we got TLE on big test cases.
We can improve this with 2 ways:
Memoization
tabulation .
*/
