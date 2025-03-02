package _D

/*
Write an efficient algorithm that searches for a value target in an 'm x n' integer matrix 'matrix'.
This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.

Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
Output: true

Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
Output: false


Constraints:

m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-10^9 <= matrix[i][j] <= 10^9
All the integers in each row are sorted in ascending order.
All the integers in each column are sorted in ascending order.
-10^9 <= target <= 10^9

[1,4,7,11,15],
[2,5,8,12,19],
[3,6,9,16,22],
[10,13,14,17,24],
[18,21,23,26,30]

   1. Brute force:
      Iterate for each item and compare with target:
      tc: n*m
      sc: 1
   2. Iterate each row and do BS in each row:
    tc: n*log(m)
    sc: 1
     Is this acceptable? Yes, but can we optimize using sorted key constraint?
   3.
    If we take rigt-top element then the row and cols are sorted.
    If we take bottom-left num, then the row and col is also sorted.
    So how can we use them?
    Lets start from right-top element.
    We can take mid and compare with target if:
    row=0
    col=len(matrix[0])-1
    for row<=len(martix)-1 && col>=0:
    mid > target
        the target never be on the rigth, in standing column, sow we need decrease column
        col=col--
    mid < target
        the target never be on the left, in standing row, sow we need icnrease row
        row++
    if == return target
    return -1
    tc: n+m
    sc: 1
    lets try implement:

*/
/*
[1,4,7,11,15],
[2,5,8,12,19],
[3,6,9,16,22],
[10,13,14,17,24],
[18,21,23,26,30]
  target = 18

 0 4

*/

func searchMatrixII(matrix [][]int, target int) bool {
	var n = len(matrix) - 1
	var m = len(matrix[0]) - 1

	var row = 0
	var col = m

	for row <= n && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target && row <= n {
			row++
		} else if col >= 0 {
			col--
		}
	}
	return false
}
