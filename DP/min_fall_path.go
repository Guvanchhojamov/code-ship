package dp

/*
931. Minimum Falling Path Sum.
Given an nxn array of integers matrix, return the minimum sum of any falling path through matrix.

A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right. Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).

Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
Output: 13.
Explanation: There are two falling paths with a minimum sum as shown.

[
[2,1,3],
[6,5,4],
[7,8,9]
]

We can use graph dp for this problem,
Can start from each cell in 0-th row, run helper functio.
helper function is:
BFS+PQ- find shortest path algorithm.
return min at the end.
Tc: (n*nlogn*n) * 3
sc: (n*n) for pq and recursion stack.
Other is we can use recursion bottom- up approach,and, then run for
each cell recursion approach.
But it will be ery slow.
TC: 3*O(2^n*n).
Sc: 3*O(n*n)
We can optimize it using, memoization, but. Its slow anyway.
So we can optimize this usign tablulation.

We need base case, define base case.
Define how we calculate next case.

*/

func minFallingPathSum(matrix [][]int) int {
	// edge cases
	n := len(matrix)
	dp := make([][]int, n)
	for j := range matrix[0] {
		dp[j] = make([]int, n)
		dp[0][j] = matrix[0][j]
	}
	//fmt.Println(dp)

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = matrix[i-1][j]
			if j > 0 && j < n-1 {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1])
			} else if j > 0 {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j-1], dp[i-1][j])
			} else if j < n-1 {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j], dp[i-1][j+1])
			}
		}
	}
	//  fmt.Println(dp)
	res := dp[n-1][0]
	for i := range dp[n-1] {
		res = min(res, dp[n-1][i])
	}
	return res
}

/*
optimized 1d memory.
*/
func minFallingPathSumMemory(matrix [][]int) int {
	// edge cases
	n := len(matrix)
	prev := make([]int, n)
	copy(prev, matrix[0])

	for i := 1; i < n; i++ {
		curr := make([]int, n)
		for j := 0; j < n; j++ {
			curr[j] = matrix[i-1][j]
			if j > 0 && j < n-1 {
				curr[j] = matrix[i][j] + min(prev[j-1], prev[j], prev[j+1])
			} else if j > 0 {
				curr[j] = matrix[i][j] + min(prev[j-1], prev[j])
			} else if j < n-1 {
				curr[j] = matrix[i][j] + min(prev[j], prev[j+1])
			}
		}
		prev = curr
	}

	res := prev[0]
	for i := range prev {
		res = min(res, prev[i])
	}
	return res
}
