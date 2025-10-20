package binarytree

/*
102. Binary Tree Level Order Traversal

Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000

*/

/*
 ok, we need to return reuslt goruped level by level traversal.
  how we can take nums level by level?
     1
   2   3
  4 5 6 7
 8
 can we achive this with recursion?
	start from root, since root is level-0, and always 1 item.
- take root, print val.
- then go to left, print val.
- then go to right, print val.
- go until node is null.
	do this cyrcle again.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res = make([][]int, 0)
	var q = make([]*TreeNode, 0)
	var levelNums = make([]int, 0)

	q = append(q, root)

	for len(q) > 0 {
		x := len(q)
		for x > 0 {
			node := q[0]
			q = q[1:]
			if node != nil {
				q = append(q, node.Left, node.Right)
				levelNums = append(levelNums, node.Val)
			}
			x--
		}
		if len(levelNums) > 0 {
			res = append(res, levelNums)
		}
		levelNums = []int{}
	}
	return res
}

/*
start with root level=0
go left, deeper increasing level.
then go right, checking if res[level] already not exisits, create empty[], else add to it.
return result at the end.
*/
func levelOrderWithDFS(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res = make([][]int, 0)

	bfs_dfs(&res, root, 0)
	return res
}

func bfs_dfs(res *[][]int, node *TreeNode, level int) {
	if node == nil {
		return
	}
	x := *res
	if level >= len(x) {
		*res = append(*res, []int{node.Val})
	} else if node != nil && level >= 0 {
		x[level] = append(x[level], node.Val)
	}
	bfs_dfs(res, node.Left, level+1)
	bfs_dfs(res, node.Right, level+1)
}

/*

    3
   / \
  9  20
     / \
    15  7
res = []
*/
