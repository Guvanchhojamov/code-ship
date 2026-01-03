package graphs

func (d *DisjointSet) find(node int) int {
	if d.parents[node] == node {
		return node
	}
	ultParent := d.find(d.parents[node])
	d.parents[node] = ultParent
	return ultParent
}

func (d *DisjointSet) union(u, v int) {
	up := d.find(u)
	vp := d.find(v)
	if up == vp {
		return
	}

	if d.ranks[up] > d.ranks[vp] {
		d.parents[vp] = up
	} else if d.ranks[up] < d.ranks[vp] {
		d.parents[up] = vp
	} else {
		d.parents[vp] = up
		d.ranks[up] += 1
	}
}

func NewDS(n int) *DisjointSet {
	p := make([]int, n)
	r := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DisjointSet{parents: p, ranks: r}
}

func numberofIslandsII(n, m int, queries [][]int) []int {
	ds := NewDS(n * m)
	result := []int{}
	islands := 0
	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, m)
	}
	matrix := visited
	for _, q := range queries {
		count := calculateIslands(q[0], q[1], n, m, &matrix, &visited, &islands, ds)
		result = append(result, count)
	}
	return result
}

// complete query and return how many islands in matrix..
func calculateIslands(r, c, n, m int, grid [][]bool, islands *int, ds *DisjointSet) int {
	if grid[r][c] {
		return *islands
	}

	grid[r][c] = true
	*islands += 1

	nodeNum := r*m + c
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range dirs {
		nr, nc := r+d[0], c+d[1]

		if nr >= 0 && nr < n && nc >= 0 && nc < m && grid[nr][nc] {
			adjNodeNum := nr*m + nc
			if ds.find(nodeNum) != ds.find(adjNodeNum) {
				ds.union(nodeNum, adjNodeNum)
				*islands -= 1
			}
		}
	}
	return *islands
}

/*
Number of Islands - II
Problem:
You are given an n, m which means the row and column of the 2D matrix,
and an array of size k denoting the number of operations.
Matrix elements are 0 if there is water or 1 if there is land.
Originally, the 2D matrix is all 0 which means there is no land in the matrix.
The array has k operator(s) and each operator has two integers A[i][0], A[i][1] means that you can change the cell
 matrix[A[i][0]][A[i][1]]  from sea to island.
 Return how many islands are there in the matrix after each operation.
You need to return an array of size k.

Note: An island means a group of 1s such that they share a common side.
   in 4 directions.

Example 1:
Input Format: n = 4 m = 5 k = 4 A = {{1,1},{0,1},{3,3},{3,4}}
Output: 1 1 2 2
Explanation: The following illustration is the representation of the operation:

Example 2:
Input Format: n = 4 m = 5 k = 12 A = {{0,0},{0,0},{1,1},{1,0},{0,1},{0,3},{1,3},{0,4}, {3,2}, {2,2},{1,2}, {0,2}}
Output: 1 1 2 1 1 2 2 2 3 3 1 1
Explanation: If we follow the process like in example 1, we will get the above result.

We need to find islands count after each query.
    Each query replaces water with land in [i][j] cell.
 it can affect other islands too.
 if cell is island and has islands in 4 directions they are connected and counts as 1 island.

Is it grapph?  yes.
Shortest path? no.
connections related? yes.
multicomponent? yes.
chacnhes state on each time? yes.
When graph realted with connections and dymaically changes, it comes up with Union find solution.
  each cell - is node.
  check 4 directional connection for each query.
  added 1 island cell on each node..
Approach -1 : Brute force DFS.
  after each each query we count disconnected islands, components with DFS. or BFS.
  - run DFS or BFS after each query added. fr4om scrach.
  - count components with DFS after each query.
 TC: K*(N*M)
 SC: (N*M)

Union find appraoch:
    union - if new cell is nei of prev island.
    find - is neighbours are in same island.
Same island means they has same ultimate prent..
   each cell is node.
we need node numbers. how we can do this?
nodes - 0..n - is to n.
after
nodes - (n+1)+col.
totalCount
if cell has nei island is 1.
 then union them
    union(cellnode, neinode)
for each new cell
   if its visited before, totalCount not changed.
if its not visited before, increment totalCount. as new island.
    check for cell neightbours if :
        they not visited and 1, then decrease totalCount -1.
        and union with this neighbour.
*/
