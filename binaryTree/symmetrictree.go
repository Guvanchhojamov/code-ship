package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}
func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) &&
		check(left.Right, right.Left)
}

/*
We need to check symmetric or not the binary tree..
What is symmetric?
The tree must be, in reversed order of left or reversed of right.
symmetric.
How we can check, root is check point.
Only root can be independend from left and right parts.
because right and left both come to root and merging.
So, for this we need to compare root->left vs root-right. Right?
What we use to traverse? BFS or DFS?
I think we need to use, dfs with preorder traversal.
1. start with root.left -> do DFS preorder traversal.
2. start with root.right -> do DFS preorder traversal.
    start both in same time.
    check 1.left != 2.right - opposide part.
 if not same return false.
 if false not returned, then return true at end.
  tc: n/2 +n/2 = O(N)
  sc: O(n) - for recursion calls.
Other ways?
1. use BFS and check q[0..n/2] == q[n...n/2] - on each level.
    if lef of parts not same we can return false, in other way.
    TC: N
    sc: len(max_level, height)
2. Can we do it with DFS but without recursion?
   us dfs preorder traversal and keep left and right stack.
   pop of them and compare.

*/
/*
 follow up approach

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetricFollow(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var q = make([]*TreeNode, 0)
	q = append(q, root.Left, root.Right)
	for len(q) > 0 {
		left := q[0]
		right := q[1]
		q = q[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		q = append(q, left.Left, right.Right)
		q = append(q, left.Right, right.Left)
	}
	return true
}
