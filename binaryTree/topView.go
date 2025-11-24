package binarytree

import "math"

/*
Problem Statement:
Given a Binary Tree,
return its Top View.
 The Top View of a Binary Tree is the set of nodes visible when we see the tree from the top.
*/
/*
 ok to solve this problem we need to see from top.
 so, we already use this method.
 What nodes we see, when watch from top?
 The first node on each level, or ,the most left or right node on each level.
 Other nodes we dont see. Or the deeper nodes we dont see.
 so, we iterate each level, take first node, add to map, continue.
 next time when we come to this node we chech is node seen in this level before if yes we just dont add this in to map,
    ignore and handle next node.
 What type of traversal we use? BFS, DFS.
 we can use both, i think, but BFS with level order traversal would be more clean in this problem.
  - create queue, create seen map, with node+level, key - boolean.
  - iterate each node add to queue left and right child.
  - until queue is empty handle all nodes.
  - if node is seen map ignore and continue
  - if not in seen map, then add to map.
  - keep max and min level vals to generate response in given order.
  tc: N
  sc: N - for map.
  some clarify notes:
    - is root can be the nil ?  check to root nil

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	Node  *TreeNode
	Level int
}

func topView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var q = make([]Queue, 0)
	var seen = make(map[int]int)
	var maxlevel, minlevel = math.MinInt, math.MaxInt
	var level = 0
	q = append(q, Queue{root, 0})
	for level < len(q) {
		item := q[level]
		if _, ok := seen[item.Level]; !ok {
			seen[item.Level] = item.Node.Val
			maxlevel = max(maxlevel, item.Level)
			minlevel = min(minlevel, item.Level)
		}
		if item.Node.Left != nil {
			q = append(q, Queue{item.Node.Left, item.Level - 1})
		}
		if item.Node.Right != nil {
			q = append(q, Queue{item.Node.Right, item.Level + 1})
		}
		level++
	}
	var result = make([]int, 0)
	for i := minlevel; i <= maxlevel; i++ {
		result = append(result, seen[i])
	}
	return result
}
