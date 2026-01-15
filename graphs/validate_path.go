package graphs

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	// bild edges list.// we can use map or LL here.
	graph := map[int][]int{}
	for _, edge := range edges {
		u := edge[0]
		v := edge[1] // we can optimize here also comparing u,v with source, destination.
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u) // since it is birirectional graph.
	}

	//3.BFS level order solution
	visited := map[int]bool{}
	q := []int{}
	q = append(q, source)
	visited[source] = true

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == destination {
			return true
		}

		adj := graph[node]
		for _, vc := range adj {
			if !visited[vc] {
				visited[vc] = true
				q = append(q, vc)
			}
		}
	}
	return false

	//2. DFS interative solution
	// stack:=[]int{}
	// visited:=map[int]bool{}

	// stack = append(stack,source)
	// visited[source] = true

	// for len(stack)> 0 {
	//     node:=stack[len(stack)-1]
	//     stack = stack[:len(stack)-1]
	//     if node == destination {
	//         return true
	//     }
	//     adj:=graph[node]
	//     for i:=len(adj)-1; i>=0;i-- {
	//         vc:=adj[i]
	//         if !visited[vc] {
	//             visited[vc] = true
	//             stack = append(stack, vc)
	//         }
	//     }
	// }

	// return false

	// 1. DFS with recursion.
	// var dfs func(node int)bool
	// var visited = map[int]bool{}
	// // visited[0] = true
	// dfs = func (node int) bool {
	//     if node == destination {
	//         return true
	//     }
	//     visited[node] = true
	//     adj:=graph[node]
	//     for _,vx:=range adj {
	//         if !visited[vx] {
	//             if dfs(vx) {
	//                 return true
	//             }
	//         }
	//     }
	//     return false
	// }
	// return dfs(source)
}

/*
we are given n vertices.
0 indexed grapth, to n-1.
- all nodes are coneected, it can be components in map.
- no edge directed itself. it means no loops? -not it can be.
- undirected map.
given source and destination.
    return is we have path from source to desitnation.

*/
/*
Input: n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
  what does mean if there has path?
  After we iterate all nodes, one visited node must be destination.
  In that way we have path to dest. otherwise we dont have path to destination.
  How we can iterate all nodes, using BFS or DFS.
  Anyway in worth case we need to see all nodes, no it does not matter what algorythm we use.. it there no much diff with DFS and BFS.
    1. DFS with recursive solution.
    2. DFS with iterative solutoin.
    4. BFS with Dequeu solution.

*/
