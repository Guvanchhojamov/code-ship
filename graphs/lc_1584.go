package graphs

import (
	"container/heap"
	"slices"
)

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

type DisjointSet struct {
	parents, sizes, ranks []int
}

func (d *DisjointSet) unionBySize(u, v int) {
	up := d.findParent(u)
	vp := d.findParent(v)

	if up == vp {
		return
	}
	if d.sizes[vp] > d.sizes[up] {
		d.parents[up] = vp
		d.sizes[vp] += d.sizes[up]
	} else {
		d.parents[vp] = up
		d.sizes[up] += d.sizes[vp]
	}
}

func (d *DisjointSet) unionByRank(u, v int) {
	u_p := d.findParent(u)
	v_p := d.findParent(v)
	if u_p == v_p {
		return
	}
	if d.ranks[v_p] > d.ranks[u_p] {
		d.parents[u_p] = v_p
	} else if d.ranks[u_p] > d.ranks[v_p] {
		d.parents[v_p] = u_p
	} else {
		d.parents[v_p] = u_p
		d.ranks[u_p] += 1 // increment only when equal rank
	}
}

func (d *DisjointSet) findParent(node int) int {
	if d.parents[node] == node {
		return node
	}
	ult_p := d.findParent(d.parents[node])
	d.parents[node] = ult_p
	return ult_p
}

func NewDisjointSet(n int) DisjointSet {
	ds := DisjointSet{
		parents: make([]int, n),
		sizes:   make([]int, n),
		ranks:   make([]int, n),
	}
	for i := range ds.parents {
		ds.parents[i] = i
		ds.sizes[i] = 1
	}
	return ds
}

type Edge struct {
	from, to, cost int
}

func minCostConnectPointsWithUnion(points [][]int) int {
	edges := createEdges(points)
	slices.SortFunc(edges, func(a, b Edge) int {
		if a.cost > b.cost {
			return 1
		}
		return -1
	})

	//fmt.Println("edges", edges)

	result := 0
	edgesCount := 0
	ds := NewDisjointSet(len(points))
	for _, ed := range edges {
		if edgesCount == len(points)-1 {
			break
		}
		if ds.findParent(ed.from) != ds.findParent(ed.to) {
			ds.unionByRank(ed.from, ed.to)
			result += ed.cost
			edgesCount++
		}
	}
	return result
}

func createEdges(points [][]int) []Edge {
	edges := []Edge{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges,
				Edge{
					i, j, abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1]),
				})
		}
	}
	return edges
}

func abs(x int) int {
	if x < 0 {
		return x * (-1)
	}
	return x
}

/*
 this problem is MST, minimum spanning tree, need to return mincost to connect all nodes.
 Before we solve this with prims algo.
 Now we try to solve this with Kruskals+Union find algo.
Approach:
    Since we need min cost, sort edges by cost:
        {from,to,cost}
    after start iterate one by one:
    take node, if from,to is in one path, in same set
    just ignore, because they aready connected.
    if they are not in same set then union() them into one set.
    and increment count of connected edges. They must be exactly V-1.
    and add cost to response.
    why it works ?
1. we take only 1 edge for connect.
    if they already connected we just skip...
2. we take min cost edge from all possibiliteis.
TC: VLogV - sort + ~V(1) - union find
SC: E - edges arr. + V union find..
*/
