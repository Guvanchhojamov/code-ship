package graphs

import (
	"container/heap"
	"math"
)

// type Item struct {
// 	to   byte
// 	cost int
// }

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

type Node struct {
	to     byte
	weight int
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {

	n := len(source)
	graph := make(map[byte][]Node)

	for i := 0; i < len(original); i++ {
		u, v, wt := original[i], changed[i], cost[i]
		graph[u] = append(graph[u], Node{v, wt})
	}
	// fmt.Println(graph, n)
	totalMinCost := 0
	for i := 0; i < n; i++ {
		if source[i] == target[i] {
			continue
		}
		pathMinCost := findPath(n, source[i], target[i], graph)
		if pathMinCost == -1 {
			return -1
		}
		totalMinCost += pathMinCost
	}

	return int64(totalMinCost)
}

func findPath(n int, src, dest byte, graph map[byte][]Node) int {

	dist := make(map[byte]int)
	for i := byte('a'); i <= byte('z'); i++ {
		dist[i] = math.MaxInt
	}

	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, &Item{src, 0})
	dist[src] = 0

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		u, cost := curr.to, curr.cost

		if cost > dist[u] {
			continue
		}
		if u == dest {
			return cost
		}

		for _, nei := range graph[u] {
			newCost := cost + nei.weight
			if newCost < dist[nei.to] {
				dist[nei.to] = newCost
				heap.Push(pq, &Item{nei.to, newCost})
			}
		}
	}

	return -1
}

/*
 We have source, target words.
  We need to get tartget from source. using.
  original and change arays.
  we can pick char for change from source, if its in original, we take it from original. and cahnge with change[i]
  if this exists, and for this we use cost[i].
So we need to get target in minimum cost.
If its impossible return -1.
ok, from descripion i feel likw its graph problem.
Because we have source, we have nodes, edges
original -> changed and wight for each edge, cost array.
We need greate adj list somehow, and try get trarget using this list.
ok, sine we need minimum we need to do this
 BFS+Dijkstra+MinHeap.
OK, we need answer some questions.
- How we convert it to adj list?
- How we iterate over source ?
- how we know we got target or no?
Ok, maybe its DP problem, with trying all possibilities,
but i will stay and try graph first.
Item:
    to, cost
PQ []Item.

Node:
    to, weight

1. Create graph:
    map[int][]Node
    key- node; value is adj list.
2. create destination arr also.
3. iterate over source and target
    until pq.len > 0 and s[i] != t[j]:
        we do BFS sohortest path.
    update dest arr.
    when we reach target[n-1]:
        return cost.
4. if at the end:
    if t[n-1] is maxint return -1.
is source.eln == target.len ? Yes.
min len for them is, 1.

*/
