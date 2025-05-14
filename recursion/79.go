package recursion

/*

79. Word Search
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.



Example 1:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
Example 2:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true
Example 3:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false


Constraints:

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board and word consists of only lowercase and uppercase English letters.
*/

package main

import (
"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}

func exist(board [][]byte, word string) bool {
	for r:=0;r<len(board);r++{
		for c:=0;c<len(board[0]);c++ {
			if dfs(board, word,r,c,0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, row, col, k int) bool{
	// base case
	if k == len(word) {
		return true
	}

	if row < 0 ||
		col < 0 ||
		row >= len(board[0]) ||
		col >= len(board) ||
		seen[key] ||
		board[row][col] != word[k] ||
		board[row][col] == '0' {
		return false
	}

	original:=board[row][col]
	board[row][colr] = '0'
	if dfs(board,word,row-1,col,k+1)||
		dfs(board,word,row,col-1,k+1)||
		dfs(board,word,row+1,col,k+1)||
		dfs(board,word,row,col+1,k+1) {
		return true
	}
	board[row][col] = original
	return false
}

/*
  for each cell call dfs depp first search, to check all possible ways.
    base cases:
        n > board || m > board[0] then return false.
        if k == len(s):
            return true
    recursion case:
        if isSafe && val == s[k]:
            dfs()

 for i to board:
    if dfs() {
        return true
    }
 retrn false.
*/


