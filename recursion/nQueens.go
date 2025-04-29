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
