package binarytree

import "math"

type QueueBT struct {
	node *TreeNode
	col  int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var q = make([]QueueBT, 0)
	var res = make([][]int, 0)
	var mp = make(map[int][]int, 0)
	var maxVal, minVal = math.MinInt, math.MaxInt
	q = append(q, QueueBT{root, 0})

	var bfs = func() {
		// todo implement in there bfs levelorder traversal
		for len(q) > 0 {
			front := q[0]
			q = q[1:]
			if front.node == nil {
				continue
			}
			maxVal = max(maxVal, front.col)
			minVal = min(minVal, front.col)
			mp[front.col] = append(mp[front.col], front.node.Val)
			q = append(q, QueueBT{front.node.Left, front.col - 1})
			q = append(q, QueueBT{front.node.Right, front.col + 1})
		}
	}

	bfs() // trigger bfs method here.

	for i := minVal; i <= maxVal; i++ {
		res = append(res, mp[i])
	}
	return res
}

/*
  ok lets try implement bfs traversal with queue.
  assume we not allowed change TreeNode structure.
  We create Dequeue simple in golang.
  and pop from front, and add to back.
  until queue is empty.
  We keep track the max and minvalues.
*/

// func verticalTraversal(root *TreeNode) [][]int{
//     if root == nil {
//         return [][]int{}
//     }
//     var mp = make(map[int][]int, 0)
//     var col = 0
//     var minval int
//     var res = make([][]int, 0)
//     mp[col] = append(mp[col], root.Val)
//     mp = dfs(node.Left, col-1,mp, &minval)

//     i:=minval
//     for {
//         if val, ok := mp[i]; !ok {
//             break
//         }
//         res = append(res, val)
//     }
//     return res
// }

// func dfs(node *TreeNode, col int, mp map[int][]int, minval *int) map[int]int{
//     if node == nil {
//         return
//     }
//     *minval = min(*minval, col)
//     mp[col] = append(mp[col], node.Val)
//     mp = dfs(node.Left, col-1, mp, minval)
//     mp = dfs(node.Right, col+1, mp, minval)
//     return mp
// }

// /*
// Vertical Traversal of a Binary Tree
// Last Updated : 07 Oct, 2025
// Given the root of a binary Tree,
// Find its vertical traversal starting from the leftmost level to the rightmost level.

// Note: If multiple nodes are at the same horizontal distance from the root and on the same level,
//  they should be printed in the order they appear in a level-order traversal (top-to-bottom, left-to-right).

// Here, Horizontal distance is calculated from the root to a specific node by counting how many
// times we move left or right along the unique path from the root to that node.

// The formula for Horizontal distance from the root is given by:

// Horizontal Distance = Number of right moves − Number of left moves in the path from the root to that node.

// */

// /*
// Input: root = [1,2,3,null,4,5,null]

// Output: [[2],[1,4,5],[3]]
//  ok to solve this problem we need to divide each vertical level to the parts.
//  so, we need each level elements, and defined columns for each node.
//     how we define column?
//     we keep column value, at root node it starts from 0, since root is in the middle, if we see tree like horizontally
//   if we go to left, it is one level up, so we decrease column val.
//   if we go to the right we incerease column value we, like we go down. on horizontal manner.
//   so, we create a map of node vals, map of lists,
//     key - is column, and val - list[of node vals]
//    we define column for each node and add this node val, to defined column list in map.
// what algorythm we need to use?
//     DFS or BFS,
//   The additional question says, if nodes are in the same column and in the same row, we need to keep order of the nodes,
//     from top -> bottom and from left to right.
//     so, if we use DFS?  witch traversal we eneed to use?
//         Preorder traversal.
//             node, left, right.
//         ok, lets try do it in, DFS with preorder traversal.

// */
