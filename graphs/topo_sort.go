package graphs

import "fmt"

// kans algo
func topoSort(V int, adj [][]int) []int {
	// calculate indegrees
	indegrees := []int{}
	for u := 0; u < V; u++ {
		for _, v := range adj[u] {
			indegrees[v]++
		}
	}

	q := []int{}
	for node, ins := range indegrees {
		if ins == 0 {
			q = append(q, node)
		}
	}
	fmt.Println(indegrees, q)
	ans := []int{}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, nei := range adj[node] {
			indegrees[nei]--
			if indegrees[nei] == 0 {
				q = append(q, nei)
			}
		}
		ans = append(ans, node)
	}
	// Follow up How to detect cyrcle?
	// at the end if we dont process all nodes, ex: any indegrees[i] == 1 we got
	// circle. All nodes must be 0 val. we can check it with counting. or queue
	// status at the beginnig.
	// if after counting indegrees, and adding 0 indegree nodes to q, if q is empty
	// we got cyrcle all nodes depends at least one node, and this is cyrcle..

	return ans
}

// dfs approach
// process all adj nodes.
// add to stack
// func topoSort(V int, adj [][]int) []int {
//     var visited = map[int]bool{}
// 	var stack = []int{}
// 	for i:=0;i<V;i++{
// 		if visited[i] == false {
// 			visited[i] = true
// 			dfs(i, adj, visited, &stack)
// 		}
// 	}
// 	//fmt.Println(stack)
// 	ans:=[]int{}
// 	for len(stack) > 0 {
// 		ans = append(ans,stack[len(stack)-1])
// 		stack = stack[:len(stack)-1]
// 	}
// 	return ans
// }

// func dfs(node int, adj [][]int, visited map[int]bool, stack *[]int) {
// 	visited[node] = true
// 	for _,nei:=range adj[node] {
// 		if visited[nei] == false {
// 			dfs(nei, adj, visited,stack)
// 		}
// 	}
// 	*stack = append(*stack, node)
// 	return
// }

/*
Input: V = 6,adj=[ [ ], [ ], [3], [1], [0,1], [0,2] ]
 When its about ordering, scedule and DAG, acyclic grapth.
 We need to keep in mind it is the topo sort graph problem.
 Because in topo sort the order is matters.
	And the graph can not contain cyrcles.
	a->b->c
	lets se this example:
	b<-a - depends a. before b me must take a.
	c<-b - depends b. before b we must take c.
  if we need print this in topo order:
  a,b,c - because we need take  first independent node. its a.
  then process it and then take next independent it is B because we processed a.
  then take c.
  a,b,c
  but if node has more than 1 deps. we must process all of them.
  To solve topo sort we have 2 ways:
	- Kans alogorythm.
	- BFS algo manual controlling.
  Kans algo easier to implement and understand.
1. DFS approach:
	- process all dependencies of node, (all not visited adj nodes) and set visited.
	- after processing it means we processed all nodes wich depends from our node.
	- add node to stack.
	- process others until end, using DFS and keeping visited array.
	- pop all nodes from stack. and this is our topo sort.
TC: N, N-stack.
SC: N - rec. call. N-stack.
It is harder some times, to keep track correctly recursion calss, and process all nodes correcty.

2. Kans algo approach:
	Kans. algo uses BFS approach, in under the hood. So we use queue
	apporach is indegrees count, if 0 indegrees it means we a free so we can be added to result.
	- compute and store in map node deps. it means adj nodes count. key-node; val-count of deps. (adj)
	- add all nodes with inderee = 0 to queue.they have no dependencies.
	- process the queue:
		take adj of curr.node; decrease -1 dep for adj node in map.
		if adj deps == 0 add it to queue to process.
	- then add curr node.val to response.
	Repeat until q is empty.
*/

func XorY(x, y int) bool {
	if x == y {
		return true
	} else if y == x {
		return true
	}
	if x != y {
		return false
	} else {
		return true
	}
}
