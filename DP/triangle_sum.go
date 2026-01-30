package dp

/*
Given a triangle array, return the minimum path sum from top to bottom.

For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.


Example 1:

Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
2
3 4
6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).

Input:
triangle  [][]int
Output:
Int min path sum.
We need to find path sum from to to bottom.
Can move DOWN and Diogonally.

Approach -1:
Can we use grah algo there?
Start from 0,0 - because its only 1 element.
Go until n-1 row. When we reach n-1 row, we got our path.
Use find short path algorithm, use BFS + PQ approach, visit bottom bottom+1 neighbour for each cell in row.
TC: ((N*m/2)log(n*m)/2)
SC: O((n*m)/2)
Any other approach?
	We can check recursively, moving greedily.
IN there we can use bottom up approach but, we have n elements in bottom.
It will take N start points, n combinations, so in top we have 1 element so we start from him.
Approach-2:
// base case:
	if i==n-1 {return grid[i][j] }
//recursion case:
bottom:=fn(i+1,j)
diogonal:=fn(i+1,j+1)
	return min(bottom, diogonal)
tc: 2^(len(grid[n-1])) - last len is max len.
sc: len(grid[n-1]) -f stack call.
How can we improve TC and space?

using a top to bottom approach.
create an empty dp [][]int triangle.
define dp[0][0] = grid[0][0]
and start from 1 th node.
compare prev+bottom, prev+diagonal.
Take and sum min
Update dp and continue until n-1 row.
m:=len(grid[n-1])
TC: O(N)
SC: O(N)

Ok,we can start from n-2 node and check above, to avoid edge cases,
Because we czech bottom and right diogonal above and its always in boundaries.

*/

func minimumTotal(triangle [][]int) int {
	// some edge cases
	n := len(triangle)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	for j := range triangle[n-1] {
		dp[n-1][j] = triangle[n-1][j]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
		}
	}
	//fmt.Println(dp)
	return dp[0][0]
}
