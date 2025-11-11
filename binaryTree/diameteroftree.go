package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var maxi = 0
	findMax(root, &maxi)
	return maxi

}

// to do this taske we use postorder traversal.
// we need calculate heigh before returning
// and at the end return maxi +1 for current subtree.
func findMax(node *TreeNode, maxi *int) int {
	if node == nil {
		return 0
	}
	lh := findMax(node.Left, maxi)
	rh := findMax(node.Right, maxi)
	*maxi = max(*maxi, lh+rh)
	return 1 + max(lh, rh)
}

// we can solve this problem using gloabal variable to store max diameter also.
// but in this case we need to restore global variable after each call.
