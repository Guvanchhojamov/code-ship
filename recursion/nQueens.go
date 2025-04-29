package recursion

/*
51. N-Queens
The n-queens puzzle is the problem of placing n queens on an n x n chessboard
such that no two queens attack each other.

Given an integer n,
return all distinct solutions to the n-queens puzzle.
You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens'
placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

Example 1:


Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above

Example 2:
Input: n = 1
Output: [["Q"]]

Constraints:

1 <= n <= 9
*/
/*
To solve this problem we need use recursion, and with this recursion.
need to check all possibilities.
start from 0,0 index paste Q in there and check is it possible paste in the next col.
next col and etc.
Start with 0,0 col and iterate over end n*n array.
 we need helper function isSafe with says can we paste in next col or not.
if not return call back and remove Q from prev row,col pasted place.
*/

func solveNQueens(n int) [][]string {
	var mainBoard = make([][][]rune, 0)
	var tmpBoard = make([][]rune, 0)
	for i := 0; i < n; i++ {
		row := []rune{}
		for j := 0; j < n; j++ {
			row = append(row, '.')
		}
		tmpBoard = append(tmpBoard, row)
	}
	fnQueen(&mainBoard, tmpBoard, 0)
	r := generateResult(mainBoard)
	return r
}

func fnQueen(mainBoard *[][][]rune, tmpBoard [][]rune, col int) {
	// base case
	if col == len(tmpBoard) {
		var newBoard = make([][]rune, len(tmpBoard))
		for i := range tmpBoard {
			newBoard[i] = make([]rune, len(tmpBoard[i]))
			copy(newBoard[i], tmpBoard[i])
		}
		*mainBoard = append(*mainBoard, newBoard)
		return
	}

	// recursion case. for each row element in this column we need to check
	for row := 0; row < len(tmpBoard); row++ {
		if isSafe(tmpBoard, col, row) {
			tmpBoard[row][col] = 'Q'
			fnQueen(mainBoard, tmpBoard, col+1)
			tmpBoard[row][col] = '.'
		}
	}

	return
}

// we need isSafe helper function to check can we paste queen in there or not.
func isSafe(board [][]rune, col, row int) bool {
	n := len(board)
	for c := 0; c < col; c++ {
		if board[row][c] == 'Q' {
			return false
		}
	}

	for r := 0; r < row; r++ {
		if board[r][col] == 'Q' {
			return false
		}
	}

	// Chep yokarky diogonaly barlamaly.
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == 'Q' {
			return false
		}
	}

	// chep ashaky diogonaly barlamaly
	for r, c := row+1, col-1; r < n && c >= 0; r, c = r+1, c-1 {
		if board[r][c] == 'Q' {
			return false
		}
	}

	return true
}

func generateResult(boards [][][]rune) [][]string {
	var res [][]string
	for _, board := range boards {
		var temp []string
		for _, row := range board {
			temp = append(temp, string(row))
		}
		res = append(res, temp)
	}
	return res
}

/*
This is the linear solution, we can optimize time, using hashmap.
but it take more space complexity.
Store queens in map, and check with o-1 time.
*/
