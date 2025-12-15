package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	// build ajd list.
	graph := map[*TreeNode][]*TreeNode{}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil && node.Right != nil {
			graph[node] = append(graph[node], node.Right)
			graph[node.Right] = append(graph[node.Right], node)
			stack = append(stack, node.Right)
		}
		if node != nil && node.Left != nil {
			graph[node] = append(graph[node], node.Left)
			graph[node.Left] = append(graph[node.Left], node)
			stack = append(stack, node.Left)
		}
	}

	res := []int{}
	visited := map[*TreeNode]bool{}
	DFS(graph, target, visited, k, &res)
	return res
}

func DFS(graph map[*TreeNode][]*TreeNode, node *TreeNode, visited map[*TreeNode]bool, k int, res *[]int) {
	if k == 0 {
		*res = append(*res, node.Val)
	}
	adj := graph[node]
	visited[node] = true

	for _, nei := range adj {
		if !visited[nei] {
			visited[nei] = true
			DFS(graph, nei, visited, k-1, res)
		}
	}

	return
}

/*
    in this porblem we need to go back, somehow.
    This is tricy part.
    But we cannot go back on binary tree.
    to go back and count exactly K elements we need some ds to store parent node.
    or more optimal algorythm.
    Where we can go to connected nodes? exactly in Graph.
    Can we buld Graph adj list from binary tree? Yes.
    So, approach:
    Assume this is undirected graph.
    we need to find all K-th neighbours from graph, well.
    - build graph adj list, iterating in DFS, or BFS. Bidirectional graph.
    - soter ajd list as map[node][]*TreeNode, all connected nodes are adj nodes.
    - in this case we build exactly what we need parent nodes go as adj nodes, so we can go back.
Start itreating on graph:
    - craete visited map.
    - iteratea adj nodes.
    - How to count?
       start iterate graph DFS from target node with distance= 0
       each time cal adj with distance+1
       when distance == k  add node to result.
    - keep track visited map to avoit infiinte loops.
 TC: N-BT; N-graph  = N
 SC: N-BT stack + N visited map, +N graph stack  = N
*/
