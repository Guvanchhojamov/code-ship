package main

import (
	"fmt"
)

/*
36. Valid Sudoku.

Determine if a 9 x 9 Sudoku board is valid.
Only the filled cells need to be validated according to the following rules:
- Each row must contain the digits 1-9 without repetition.
- Each column must contain the digits 1-9 without repetition.
- Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:
 A Sudoku board (partially filled) could be valid but is not necessarily solvable.
 Only the filled cells need to be validated according to the mentioned rules.


Example 1:


Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true

Example 2:

Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8.
Since there are two 8's in the top left 3x3 sub-box, it is invalid.


Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/

/*
  We need validate sudoku desk. Length is alys 8x8 for 0-indexed matrix.
	- validate each 3x3.
	- validate whole matrix.
    - need to validate only filled cells - this means we need skip if cell == '.'
  Bf:
	- take each 3x3 matrix and validate each row, column. 9*(3^2*3^2) tc.
	- take whole matrix and validate each row, column.  9*9.
optimal :
using hash map:
	create 3 slices with, maps:
		rows, cols, boxes.
	rows =  []key - num, val - boolean.
	cols =  []key - num, val - boolean.
	boxes = []key - num, val -bool. box_index we use formula to take index.

	we have 9 boxes with 0-3,3-6,6-9 ... indexes to calculate, in with box placed the number we use formula.
	box_index = (rowNumber/3)*3 + (colNumber/3) - Why need to define.

*/

func main() {
	//board := [][]byte{
	//	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}
	//board2 := [][]byte{
	//	{'8', '4', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '3', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}
	//fmt.Println(board2)

	borad3 := [][]byte{
		{'.', '.', '5', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '1', '4', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '9', '2', '.', '.'},
		{'5', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '3', '.'},
		{'.', '.', '.', '5', '4', '.', '.', '.', '.'},
		{'3', '.', '.', '.', '.', '.', '4', '2', '.'},
		{'.', '.', '.', '2', '7', '.', '6', '.', '.'},
	}
	//board4 := [][]byte{
	//	{'.', '.', '.', '.', '.', '.', '5', '.', '.'},
	//	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	//	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	//	{'9', '3', '.', '.', '2', '.', '4', '.', '.'},
	//	{'.', '.', '7', '.', '.', '.', '3', '.', '.'},
	//	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	//	{'.', '.', '.', '3', '4', '.', '.', '.', '.'},
	//	{'.', '.', '.', '.', '.', '3', '.', '.', '.'},
	//	{'.', '.', '.', '.', '.', '5', '2', '.', '.'},
	//}

	r := isValidSudoku(borad3)
	fmt.Println(r)
}

func isValidSudoku(board [][]byte) bool {
	var rows = make([]map[byte]bool, len(board))
	var cols = make([]map[byte]bool, len(board[0]))
	var boxes = make([]map[byte]bool, 9)

	for i, row := range board {
		for j, num := range row {
			if num == '.' {
				continue
			}
			if rows[i] == nil {
				rows[i] = map[byte]bool{}
			} else if _, ok := rows[i][num]; ok {
				return false
			}
			if cols[j] == nil {
				cols[j] = map[byte]bool{}
			} else if _, ok := cols[j][num]; ok {
				return false
			}

			boxi := (i/3)*3 + (j / 3)
			if boxes[boxi] == nil {
				boxes[boxi] = map[byte]bool{}
			} else if _, ok := boxes[boxi][num]; ok {
				return false
			}
			rows[i][num] = true
			cols[j][num] = true
			boxes[boxi][num] = true
			//fmt.Println(rows[i], cols[j], boxes[boxi], i, j, boxi)
		}
	}

	return true
}

// with helper function much complex
func isAlreadySeen(seen []map[byte]bool, idx int, number byte) bool {
	if seen[idx] == nil {
		seen[idx] = map[byte]bool{}
	}
	if seen[idx][number] {
		return true
	}
	seen[idx][number] = true

	return false
}

func isValidSudokuWithHelper(board [][]byte) bool {
	var rows = make([]map[byte]bool, len(board))
	var cols = make([]map[byte]bool, len(board[0]))
	var boxes = make([]map[byte]bool, 9)

	for i, row := range board {
		for j, num := range row {
			if num == '.' {
				continue
			}

			boxi := (i/3)*3 + (j / 3)
			if isAlreadySeen(rows, i, num) ||
				isAlreadySeen(cols, j, num) ||
				isAlreadySeen(boxes, boxi, num) {
				return false
			}
		}
	}

	return true
}
