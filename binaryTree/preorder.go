package binarytree

/*
144. Binary Tree Preorder Traversal
Given the root of a binary tree, return the preorder traversal of its nodes' values.

Example 1:

Input: root = [1,null,2,3]

Output: [1,2,3]

Explanation:

Example 2:

Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]

Output: [1,2,4,5,6,7,3,8,9]

Explanation:

Example 3:

Input: root = []

Output: []

Example 4:

Input: root = [1]

Output: [1]

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100


Follow up: Recursive solution is trivial, could you do it iteratively?
*/
/*
What is preorder traversal: root,left,right.
we must see root first, then left, then right.

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func preorderTraversal(root *TreeNode) []int {
// 	var result = make([]int, 0)
// 	if root == nil {
// 		return []int{}
// 	}
// 	fn(root, &result)
// 	return result
// }

// func fn(node *TreeNode, r *[]int) {
// 	if node == nil {
// 		return
// 	}
// 	*r = append(*r, node.Val)
// 	fn(node.Left, r)
// 	fn(node.Right, r)
// 	return
// }

func preorderTraversal(root *TreeNode) []int {
	var result = make([]int, 0)
	if root == nil {
		return []int{}
	}
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

/*
Preorder traversal: root, left, right.
root, left, right.
so, if its not nil we store datam then go to leftm then go to right.
*/

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var stack = make([]*TreeNode, 0)
	var result = make([]int, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			result = append(result, node.Val)
			stack = append(stack, node.Right, node.Left)
		}
	}
	return result
}
