package stack

/*85. Maximal Rectangle
Given a rows x cols binary matrix filled with 0's and 1's,
find the largest rectangle containing only 1's and return its area.

Example 1:

Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 6
Explanation: The maximal rectangle is shown in the above picture.
Example 2:

Input: matrix = [["0"]]
Output: 0

Example 3:
Input: matrix = [["1"]]
Output: 1

Constraints:

rows == matrix.length
cols == matrix[i].length
1 <= row, cols <= 200
matrix[i][j] is '0' or '1'.

*/
/*
 We need caclulate max rectangle in 2d arr.
To do this we need height and with. How can we finc height?
What is the height of [3,0] rectangle for example?
the height is sum of a[:3,0].
the wight is j-psei[i]-1.
So for each row we need find heights and psei().
We need precompute for prefixSum of 2d array.
and to LargeArea on this prefixSummed 2d array. because the
prefixSum 2d arr gives height for each psa[i,j] we only need calculate large area for each
ps[i,j].
1. Precompute prefixSum 2d arr.
2. calculate largeArea for each row in pfsum arr and return max.
0 1
1 0

0 1
1 1
this is egde case so if the pf[i][j] ==0 we need flush sum=0 because it breaks rectangle.
*/

func maximalRectangle(matrix [][]byte) int {
	var pfmatrix = make([][]int, len(matrix), len(matrix))
	for i := range matrix {
		row := make([]int, len(matrix[0]), len(matrix[0]))
		pfmatrix[i] = row
	}

	for i := range matrix[0] {
		sum := 0
		for j := range matrix {
			if int(matrix[j][i]-'0') == 0 {
				sum = 0
			}
			sum += int(matrix[j][i] - '0')
			pfmatrix[j][i] = sum
		}
	}
	var maxArea = 0
	for i := range pfmatrix {
		area := largestRectangleAreaOptimal(pfmatrix[i])
		maxArea = max(maxArea, area)
	}
	return maxArea
}
