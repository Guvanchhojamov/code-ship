package graphs

/*
Problem Description: Breadth-First Search (BFS) in a Graph

You are given an unweighted graph,
which may be represented either as an adjacency list or an adjacency matrix.
The graph can be directed or undirected, and it may contain cycles.
Your task is to perform a Breadth-First Search (BFS) starting from a given source vertex start.
The goal of the BFS traversal is to visit all vertices reachable from the starting vertex in level order.
*/
func BFSGraph(graph map[int][]int, start int) []int {
	res := []int{}
	visited := map[int]bool{}
	q := []int{}
	q = append(q, start)
	//visited[start] = true

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if visited[node] {
			continue
		}

		res = append(res, node)
		visited[node] = true
		adj := graph[node]

		for _, vertic := range adj {
			if !visited[vertic] {
				q = append(q, vertic)
			}
		}
	}
	return res

}

/*
assume graph is 0-indexed.
 to traverse with bfs in grapth we use bfs levelorder traversal using queue ds.
 since graph may contain cirlce we need to keep track visited map witch contains, key-node; val-boolean.
 add start node to map and queue as visited,
 then take all adj nodes of this node, iterate them level by level adding to the map and response.
 to all graph, we must use visited map to avoid cyrcle checks, all nodes must be visited only once.
 tc: O(n) - we iterate all nodes.
 sc: O(n) - for queue.
*/

/*
Problem: Depth-First Search (DFS) in a Graph
You are given a graph, which may be directed or undirected,
and may contain cycles.
The graph is represented as an adjacency list.

Your task is to perform a Depth-First Search (DFS) starting from a given source vertex start.
*/
func dfsGrapth(graphs map[int][]int, start int) []int {
	result := []int{}
	visited := map[int]bool{}
	var dfs func(node int)

	dfs = func(node int) {
		result = append(result, node)
		visited[node] = true
		adj := graphs[node]
		for _, v := range adj {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	dfs(start)
	return result
}

/*
- is graph directed? or undirected?

- is grapth may contain cyrcles?

- in this problem depth for each node, we need to use DFS with preorder traversal algorythm here.

  - result []int; visited map[int]int;
    we use recursion call for each node, to go depth.
    graph = {
    0: [1, 2],
    1: [2],
    2: [0, 3],
    3: [3]
    }

    start = 2
    0
    / \
    1 - 2
    |
    3<->3
    out = [2,0,1,3];
*/
func DFSGraphIterative(graphs map[int][]int, start int) []int {
	stack := []int{}
	result := []int{}
	visited := map[int]bool{}
	stack = append(stack, start)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node] {
			continue
		}

		result = append(result, node)
		visited[node] = true

		adj := graphs[node]
		for i := len(adj) - 1; i >= 0; i-- {
			if !visited[adj[i]] {
				stack = append(stack, adj[i])
			}
		}
	}
	return result
}

/*
 DFS iterative solution with stack. LIFO.
 - stack - to keep nodes.
 - visited - to avoid cyrcle in graph.
 - result  - to return result array.

*/
