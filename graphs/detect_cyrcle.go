package graphs

type NodeGraph struct {
	val, parent int
}

func detectCyrle(graph map[int][]int, n int) bool {
	var visited = map[int]bool{}
	for i := 1; i <= n; i++ {
		if !visited[i] {
			if detecCyrcleBFS(graph, i, visited) {
				return true
			}
			// or we can use DFS also
			if detectCyrcleDFS(graph, i, visited) {
				return true
			}
		}
	}
	return false
}

func detectCyrcleDFS(graph map[int][]int, start int, visited map[int]bool) bool {
	var dfs func(start, parent int) bool

	dfs = func(node, parent int) bool {
		visited[node] = true
		adj := graph[node]
		for _, nb := range adj {
			if !visited[nb] {
				if dfs(nb, node) {
					return true
				}
			} else if nb != parent {
				// visited neighbour not equal to parent
				return true
			}
		}
		return false
	}
	return dfs(start, -1)
}

func detecCyrcleBFS(graph map[int][]int, start int, visited map[int]bool) bool {
	var q = []NodeGraph{}

	q = append(q, NodeGraph{start, -1})
	visited[start] = true

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		adj := graph[node.val]
		for _, nb := range adj {
			if !visited[nb] {
				visited[nb] = true
				q = append(q, NodeGraph{nb, node.val})
			} else if nb != node.parent {
				return true
			}
		}
	}
	return false
}

/*
Given an undirected graph with V vertices and E edges,
check whether it contains any cycle or not.

Example 1:
Input:
V = 8, E = 7

Output:  0
Explanation: No cycle in the given graph.
Example 2:
Input:
V = 8, E = 6

Output: 1
Explanation: 4->5->6->4 is a cycle.
*/

/*
Ok we need to detect circle in graph.
What is cyrcle means in graph?
 It means if we start from somewhere, and compe to the start point,
    and this start point is not parent of current point int means we detect the cirle.
 - Assume given graph list of adjesentes. What can we use DFS or BFS,
    in any type we can come again to start point. so lets first use BFS, then try to use DFS.
- values:
    Node:
        val int
        parent int
    Queue []Node
   visited map to skip already visited nodes and detect cyrcle if it is visitor is not prev node.
   if visitor is not prev node return true.
- add start node to queue with visitor 0, or -1 since no one can visit before we start.
- then take adjecentes and visit one by one and fill queue and visited map also.
- any time if node.val in visited and node.parent is not curr node.val  then return true else false.
Adn we need to ask most important thing is it graph can be component or not?
    If given n nmber it can be graph with components otherwise we assume it is not componented,
    we must ask from interviever and define this situation..
TC: O(N)
SC: O(n) for queue.
*/
