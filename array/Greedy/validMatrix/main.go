package main

func main() {
	rowSum := []int{3, 8}
	colSum := []int{4, 7}
	restoreMatrix(rowSum, colSum)
}

/*

Example 2:

Input: rowSum = [5,7,10], colSum = [8,6,8]
Output: [[0,5,0],
         [6,1,0],
         [2,0,8]]
*/

/*
    8 6 8
  5 0 0 0
  7 0 0 0
 10 0 0 0


 1. set min(rows[i], cols[j]) for col 0
    8 6 8
  5 5 0 0
  7 3 0 0
10  0 0 0
*/

func restoreMatrix(rowSum []int, colSum []int) [][]int {

	return nil
}
