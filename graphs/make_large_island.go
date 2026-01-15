package graphs

type DisjointSet struct {
	parents, sizes []int
}

func (d *DisjointSet) unionBySize(u, v int) {
	up := d.find(u)
	vp := d.find(v)
	if up == vp {
		return
	}
	if d.sizes[up] > d.sizes[vp] {
		d.parents[vp] = up
		d.sizes[up] += d.sizes[vp]
	} else {
		d.parents[up] = vp
		d.sizes[vp] += d.sizes[up]
	}
}

func (d *DisjointSet) find(node int) int {
	if d.parents[node] == node {
		return node
	}
	ultParent := d.find(d.parents[node])
	d.parents[node] = ultParent
	return ultParent
}

func NewSet(n int) *DisjointSet {
	parents := make([]int, n)
	sizes := make([]int, n)
	for i := range parents {
		parents[i] = i
		sizes[i] = 1 // on start all node size is 1.
	}
	return &DisjointSet{
		parents: parents,
		sizes:   sizes,
	}
}

func largestIsland(grid [][]int) int {
	// create union ds and fill, connect islands.
	// itereate again and if cell is 0 try connect 4 nei and check size
	// with maxSize
	/*
	   1,0
	   0,1
	*/
	n := len(grid)
	ds := NewSet(n * n)
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != 1 {
				continue
			}
			for k := 0; k < 4; k++ {
				nr := row + dirs[k][0]
				nc := col + dirs[k][1]
				if (nr >= 0 && nr < n) &&
					(nc >= 0 && nc < n) &&
					grid[nr][nc] == 1 {
					node := (row * n) + col
					adjNode := (nr * n) + nc
					ds.unionBySize(node, adjNode)
				}
			}
		}
	}
	result := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != 0 {
				continue
			}
			uniqueParents := map[int]bool{}
			for k := 0; k < 4; k++ {
				nr := row + dirs[k][0]
				nc := col + dirs[k][1]
				if (nr >= 0 && nr < n) &&
					(nc >= 0 && nc < n) &&
					grid[nr][nc] == 1 {
					adjNode := (nr * n) + nc
					uniqueParents[ds.find(adjNode)] = true
				}
			}
			size := 0
			for p, _ := range uniqueParents {
				size += ds.sizes[p]
			}
			result = max(result, size+1) // add curr 0 as 1
		}
	}
	for _, p := range ds.parents {
		if ds.parents[p] == p {
			result = max(result, ds.sizes[p])
		}
	}
	return result
}

/*
 We are given matrix, and we need to return
 max size of land at the end.
 - allowed change only one 0 to 1
 - need return island size.

Is it related with graph? Yes.
Is shortest path? No.
Is ordered ? no
is related with connections? yes. connect 1 land and make bigger.
is graph change dynamically? yes.

So, when it is graph and changes dynamically and related with connections
we need use UnionFind or DFS/BFS brute force.
approach-1:
Use brute force with BFS to count lands size.
 - iterate over using BFS once, and define maxLandsize.
 then for each 0 make this 1 and run BFS fro whole matrix to calculate
 size of max component. Compare with previous max and update.
 - Do this for each 0 cell.
TC: N*N + (N*N)*(N*N) - in worth case.
SC: N*N.

How we can optimze with using UF?
we need size of land each time. and need effienectly calluated
for seperate components.
 - for this we can use rank by size method from UF..
 - for each not land 0. use UF union and find methods.
 - to define is new land 0->1 is connected with other big land or not
    we use find().
 - get size for this parent from sizes arr add +1 our curr land.
 - compare with maxIsland.
 at the end return maxLand.
TC: N*n*alpha(N*n) - in woth case.
SC: N*n - for uf.
any edge case? maybe. what is all matirx is 1.
 at the end UF returns correct sie from sizes[parent] arr.

*/
