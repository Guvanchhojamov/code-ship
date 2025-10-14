package binarytree

/*
94. Binary Tree Inorder Traversal
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:

Input: root = [1,null,2,3]

Output: [1,3,2]

Example 2:
Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Output: [4,2,6,5,7,1,3,9,8]


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
 What is inorder traversal?
	left,root,right.
*/

func inorderTraversal(root *TreeNode) []int {
	var result = make([]int, 0)
	if root == nil {
		return []int{}
	}
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

/*
 ok, now we need to solve this with stack algo.
*/

/*
  - Go to left and add to stack depth until reach null,
  - after reaching null:
    pop from stack, add val to res.
    go to the right.
  - go until stack is empty.
*/
func inorderTraversalStack(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}

	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		node = node.Right
	}
	return result
}
