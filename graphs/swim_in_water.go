package graphs

import "container/heap"

/*
778. Swim in Rising Water
*/

type cell struct {
	row, col, time int
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n) // all false
	}

	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, &cell{0, 0, grid[0][0]})
	maxReqTime := grid[0][0]

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*cell)
		if visited[curr.row][curr.col] {
			continue
		}
		visited[curr.row][curr.col] = true
		maxReqTime = max(maxReqTime, curr.time)

		if curr.row == n-1 && curr.col == n-1 {
			return maxReqTime
		}

		dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, d := range dirs {
			nr := curr.row + d[0]
			nc := curr.col + d[1]

			if (nr >= 0 && nr < n) &&
				(nc >= 0 && nc < n) &&
				!visited[nr][nc] {
				heap.Push(pq, &cell{nr, nc, grid[nr][nc]})
			}
		}
	}
	return maxReqTime
}

/*
 While there is greed, we don't need to create adj list.
 Need to define edge cases also.
*/

/*
 So we are given an nxn matrix. We need to reach n-1xn-1 cell with given conditions.
  - start points is 0x0.
  - dest is n-1xn-1
  - we can move 4 directionally.
conditions:
  - each cell will full with water in T time. its cell value m[i][j].
  - we can swim to cell only if T is currentTime>=adjTime.
  - We can move infinitely in 1 step.. !important.
  - curr start time is 0, we add next adj time assuming we wait X time until water fulls.
    as currTime = currTime + adjTime.


How can we solve this?
What is it?
    It looks like a graph? yes.
    Realted with shortest path? Similar but we have time there and can move faster..
    Cell - node.
    4adj is edges.
    edges are weighted, because we will wait until adj cell is full.
    Is graph, connections changes dynamically? Maybe.
If its shoretest path we can use Dijkstras algo with BFS + MinHeap.
If its is dynamic calculate something we can use UnionFind algo..
Let's take examples ...
One appraoch is, Dijsktra, BFS+MinHEap. Keeping max from popen elements from minheap.
Maximim poped number from heap is our Min required time, and our result.
TC: (N*N)-BFS and Queue pop push  * logE.
SC: N*N for BFS and N*N for Queue.
Is it ok ?
maybe. But we don use advantage of we can move faster...
Maybe we can optimze if we use this? Maybe not,
lets implement this..
*/
