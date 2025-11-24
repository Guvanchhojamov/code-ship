package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if p.Val != q.Val {
		return false
	}
	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}
	return true
}

/*
 What is same tree?
 They must have same structure and same values on each equvalent node.
 BF:
 in there we use preoreder traversal.
 on each node we need compare values,
 if they are not equal return false.
 else continue to the next node, until end.
 tc: O(min(n,m)) n and m are nodes count in each tree.
 sc: O(h) h is height of tree, recursion stack.
*/
