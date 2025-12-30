package graphs

import "container/heap"

type item struct {
	node int
	wt   int
}

type PQ []*item

func (p PQ) Len() int           { return len(p) }
func (p PQ) Less(i, j int) bool { return p[i].wt < p[j].wt } // Min-Heap
func (p PQ) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *PQ) Push(x any) {
	*p = append(*p, x.(*item))
}

func (p *PQ) Pop() any {
	old := *p
	n := len(old)
	it := old[n-1]
	*p = old[0 : n-1]
	return it
}

type edge struct {
	to int
	wt int
}

func minCostConnectPoints(points [][]int) int {
	V := len(points)
	graph := getConnectedGraph(points)
	//fmt.Println(graph)

	pq := &PQ{}
	heap.Push(pq, &item{0, 0})
	visited := make([]bool, V)

	totalSum, edgeCount := 0, 0
	for pq.Len() > 0 && edgeCount < V {
		curr := heap.Pop(pq).(*item)
		if visited[curr.node] {
			continue
		}
		visited[curr.node] = true
		totalSum += curr.wt
		edgeCount += 1

		for _, nei := range graph[curr.node] {
			if !visited[nei.to] {
				heap.Push(pq, &item{nei.to, nei.wt})
			}
		}

	}
	return totalSum

}

func getConnectedGraph(points [][]int) [][]edge {
	graph := make([][]edge, len(points))
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			u, v := i, j
			wt := abs(abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1]))
			graph[u] = append(graph[u], edge{v, wt})
			graph[v] = append(graph[v], edge{u, wt})
		}
	}
	return graph
}

func abs(x int) int {
	if x < 0 {
		return (x) * (-1)
	}
	return x
}

/*
  Given pints, we can say it like nodes.
  we need connect this points as minimum possible weight way.
  weight = abs(xi-xj)+abs(yi+yj) absolute weight of long lat of 2 points
  What we need to do?
  Since we are given points with lang lat.
  if we ssume points are nodes. and calculate weight for each edge.
  a->b weight from a->to b and store it somehow.
  it becomes, weighted undirected graph.
  - for now we assume we dont have circles.
    is can be there maybe? maybe not.
  So, how to calculate weight and store it?
   connect each point with abs weight.
   - after find min spanning tree for this connected nodes.
   using PQ, and Primes algo.

*/
