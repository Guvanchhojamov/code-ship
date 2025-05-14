package recursion

/*
This is the same with Rat in a Maze problem.

1219. Path with Maximum Gold

In a gold mine grid of size m x n,
each cell in this mine has an integer representing the amount of gold in that cell,
0 if it is empty.
Return the maximum amount of gold you can collect under the conditions:

 - Every time you are located in a cell you will collect all the gold in that cell.
 - From your position, you can walk one step to the left, right, up, or down.
 - You can't visit the same cell more than once.
 - Never visit a cell with 0 gold.
 You can start and stop collecting gold from any position in the grid that has some gold.


Example 1:

Input: grid = [[0,6,0],[5,8,7],[0,9,0]]
Output: 24
Explanation:
[[0,6,0],
 [5,8,7],
 [0,9,0]]
Path to get the maximum gold, 9 -> 8 -> 7.
Example 2:
Input: grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
Output: 28
Explanation:
[[1,0,7],
 [2,0,6],
 [3,4,5],
 [0,3,0],
 [9,0,20]]
Path to get the maximum gold, 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7.

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 15
0 <= grid[i][j] <= 100
There are at most 25 cells containing gold.
*/

/*
  We can solve this using recursion and backtracking algorithm.
	1.start from 0,0
	2. check all posiibilities in recursive func.
	3. move to 0,1 and so on until n,n.
and return max element.
tc: n*m * (4^n*m)
*/

/*
When a problem can be solved with depth-first search,
it can often also be solved with breadth-first search (BFS).

We will create a function, bfsBacktrack,
that uses a breadth-first search to find the path with the maximum gold
for a given starting cell.
We will use a queue to store the cells we need to search.
Each entry in the queue contains the coordinates of the current cell,
the gold found so far on the path, and a set storing the cells visited on this path so far.

When we pop the front cell from the queue,
we store the amount of gold found on the path so far as currGold,
and update the maxGold if the currGold is higher.

Then, if each of the four adjacent cells has gold,
is inside the matrix, and has not yet been visited,
we mark them as visited and add them to the queue with the updated gold collected.
After adding the cell to the queue, we remove it from the visited set to explore
other possible paths from this cell during backtracking.

To improve the efficiency of the solution,
we calculate the total amount of gold in the matrix before searching.
This way, if we discover a path that has the maximum possible total gold,
we can halt the search process.
Similar to the above solution, we call bfsBacktrack for every starting cell in the matrix.

*/
/*
[[0,6,0],
 [5,8,7],
 [0,9,0]]

should be solved.
*/
