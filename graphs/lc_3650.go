package graphs

import (
	"container/heap"
	"math"
)

type Item struct {
	u, cost int
}

// // Priority Queue implementation for Go
// type PQ []*Item

// func (pq PQ) Len() int           { return len(pq) }
// func (pq PQ) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
// func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
// func (pq *PQ) Push(x any)        { *pq = append(*pq, x.(*Item)) }
// func (pq *PQ) Pop() any {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	*pq = old[0 : n-1]
// 	return item
// }

// type Edge struct {
// 	to, weight int
// }

func minCost(n int, edges [][]int) int {

	// create adj list.
	graph := make(map[int][]Edge)
	for _, edge := range edges {
		u, v, wt := edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], Edge{v, wt})
		graph[v] = append(graph[v], Edge{u, 2 * wt})
	}

	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, &Item{0, 0})

	dest := make([]int, n) // 0-normal min weight, 1 - rev 2*w min weight.
	for i := range dest {
		dest[i] = math.MaxInt
	}
	dest[0] = 0

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		u, cost := curr.u, curr.cost

		if cost > dest[u] {
			continue
		}
		if u == n-1 {
			return cost
		}

		for _, nei := range graph[u] {
			newCost := cost + nei.weight
			if newCost < dest[nei.to] {
				dest[nei.to] = newCost
				heap.Push(pq, &Item{nei.to, newCost})
			}
		}
	}

	//  fmt.Println(dest)
	if dest[n-1] == math.MaxInt {
		return -1
	}
	return dest[n-1]
}

/*
 Notes:
  input:
    - n int
    - edges [][]int.
  output:
    - minCost int.

 Constraints:
    - n range ? 0<=n <= 10^5.
    - nodes are unique? Yes.
    - always directed graph? Yes.
    - Can contain ccircles? Yes or no. Probably yes.
    - Is there can be negative weights?   min wi = 0? Yes. min is 0.
    - We can change direction ui->vi; and cost will be 2*wi.
    - We can chahge swich only once.
*/
/*
  Ok, since we have directed and weighted grapth and we need to find shortest path from 0 -> n-1 node.
  We need try all possible paths and take min cost. from them .
  Since for shortest path problems we can use Dijkstra + BFS and  Priority queue algorythm, i think.
  Approach:
    - create pq prioritizing weight.
    - BFS starts from 0 th node. addign to pq first node with weihgt.
    - keep visited array also with len N. But how?
        There can be circles so we need to keep state also. For this we use 0,1,2 vals.
            0 - not visited. We can visit this node.
            1 - already in this path and someone keep this, circle immediately return -1.
            2 - visited, done, just skip.
    - keep destination arr with len N. fill with INT_MAX. Since we need min path.
    - keep loop until pq emty or we reach last node.
    - pop node from heap
    - if its last node, update with dest[n-1] and curr weight. and quit from loop.
    - compare weight with dest[node] update to min.
    - add neigthbour nodes with updated weights, and important with switch weights.
        2 * w_nei_node. Maybe this is lesser.
    - return dest[n-1] at the end.
Do we really need to check circle? Maybe no. Iterate all nodes.
At the end if dest[n-1] == INT_MAX we cannot reach end. retrun -1.
Ok sice we need rev_adj list to see where we can go on reversed.
TC: 2*N with swiches. Assume we does switch for each node. and this is answer. But its n-1 we cannot change last direction.
SC: 2*N-PQ - again same.
*/
