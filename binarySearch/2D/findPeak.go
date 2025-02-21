package _D

/*
1901. Find a Peak Element II
A peak element in a 2D grid is an element that is strictly greater than
all of its adjacent neighbors to the left, right, top, and bottom.

Given a 0-indexed m x n matrix mat
where no two adjacent cells are equal,
find any peak element mat[i][j] and
return the length 2 array [i,j].

You may assume that the entire matrix is surrounded by an outer perimeter with the value -1 in each cell.

You must write an algorithm that runs in O(m log(n)) or O(n log(m)) time.

Input: mat = [[1,4],[3,2]]
Output: [0,1]
Explanation: Both 3 and 4 are peak elements so [1,0] and [0,1] are both acceptable answers.
1 4
3 2

Input: mat = [[10,20,15],[21,30,14],[7,16,32]]
Output: [1,1]
Explanation: Both 30 and 32 are peak elements so [1,1] and [2,2] are both acceptable answers.
[10,20,15],
[21,30,14],
[7, 16,32]
1 1
2 2
*/

/*
  What is the peak element.
  Since we know how to find peak element in 1d array.
  For 2d array this logic works to, think how can we eliminate rigth or left part?

  1. Simple bruteforce is:
    Iterate for each element and find max value in matrix, this is can be our ans.
    tc: n*m
    sc: 1
  2. How can we optimize?
   Is array is sorted? no
   Is we need sorted array?
   In other words we can say,
    the peak would be max of row and col
   We start from rightmost col and go to left,
   determine right part.
   select last col,
    find max element in this column.
    then check left and rigth of this element.
    if bigger than left and rigth it is peak element we can return index.
    if left>el move colnumber to left
    if el<right move colnumber to right
    if left>el then left is greater all of el column numbers because el is max in column
    if el<right then righ is greater all of el column numbers because el is max in column.
*/
/*
[10,20,15],
[21,30,14],
[7, 16,32]
*/
/*
   0 1
 0 1 4
 1 3 2
 low  = 0
 high = 1
 mid = 0
 col =
*/
/*
 [47,30,35,8,25],
 [6,36,19,41,40],
 [24,37,13,46,5],
 [3,43,15,50,19],
 [6,15,7,25,18]
   [2,3
*/
func findPeakGrid(mat [][]int) []int {
	var n = len(mat)
	var low, high = 0, n - 1

	for low <= high {
		mid := (low + high) / 2
		col := getMaxInRow(mat, mid)
		top, bottom := 0, 0

		// cover edge cases
		if mid-1 < 0 {
			top = mat[0][col]
		} else {
			top = mat[mid-1][col]
		}
		if mid+1 > n-1 {
			bottom = mat[n-1][col]
		} else {
			bottom = mat[mid+1][col]
		}
		//  fmt.Println(top,bottom,mid,col)
		if mat[mid][col] >= top && mat[mid][col] >= bottom {
			return []int{mid, col}
		}
		if mat[mid][col] < top {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return []int{}
}

func getMaxInRow(mat [][]int, rowI int) (colI int) {
	var maxEl = mat[rowI][0]
	for i := 1; i < len(mat[0]); i++ {
		if mat[rowI][i] > maxEl {
			maxEl = mat[rowI][i]
			colI = i
		}
	}
	return
}
