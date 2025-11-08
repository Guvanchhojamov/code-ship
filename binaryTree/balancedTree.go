package binarytree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return checkHeight(root) != -1
}

func checkHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lh := checkHeight(node.Left)
	rh := checkHeight(node.Right)

	if lh == -1 || rh == -1 {
		return -1
	}

	if math.Abs(float64(lh)-float64(rh)) > 1 {
		return -1
	}

	return 1 + max(lh, rh)
}

/*
 Ok, what is balanced, height balanced binary tree?
 the diff between height.
 On every node, the diff left-childs and right-childs must be less than 0.
 if each. root:
    height(left) - height(right) > 1:
    return false : return true.
BF:
on each node we need get left nodes count,
and right nodes cound and compare.
if diff > 1 then return false.
 else contniue to the next node, until end.
 tc: n*2h, for each node we go until depth, right and left.
 sc: 1

*/
