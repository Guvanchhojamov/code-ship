package binarytree

import "math"

type Nod struct {
	node *TreeNode
	hd   int
}

func bottomView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	mini, maxi := math.MaxInt, math.MinInt
	visible := map[int]int{}
	q := []Nod{}
	q = append(q, Nod{root, 0})

	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		if item.node == nil {
			continue
		}
		visible[item.hd] = item.node.Val
		maxi = max(maxi, item.hd)
		mini = min(mini, item.hd)
		q = append(q, Nod{item.node.Left, item.hd - 1}, Nod{item.node.Right, item.hd + 1})
	}

	for i := mini; i <= maxi; i++ {
		result = append(result, visible[i])
	}

	return result
}

/*
Problem Statement:
 Given a Binary Tree, return its Bottom View.
 The Bottom View of a Binary Tree is the set of nodes visible when we see the tree from the bottom.
*/
/*
 The bottom view is, last node for each level in, horizontal manner.
 it means if we met new node in same level we can update value.
 If we go BFS levelorder traversal + level
    - create map, add key-level; val - node.val. if we go left=level-1; right=level+1
 at the end all nodes in map should be bottom view nodes.
*/
