package binarytree

/*
145. Binary Tree Postorder Traversal
Given the root of a binary tree, return the postorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [3,2,1]

Explanation:
*/
/*
 What is post order traversal? left, right, root.
 Visit in the post order, so we need withit left, then right, then root.
    1. Solve with recursive approach.
    2. Then solve with stack and sequental approach.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal2(root *TreeNode) []int {
	var res = make([]int, 0)

	if root == nil {
		return []int{}
	}
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

/*
 tc: N
 sc: N - recursion stack call.
*/

/*
 ok, now how we can solve it sequentally?
 Input: root = [1,null,2,3]
Output: [3,2,1]
   1
  *  2
    3
- start with root 1.
- go left until last left.
  and add to stack.
  add right to stack.
- then pop from stack.
-


4 2 5 3 -1 7 6 -1 9 -1 -1 8 -1 1
ok, to iterate postorder traversal we have
1. since, postorder is: left, right, root - take root in end its hard, todo so we need to use tricky.
    reverse postorder logic: left->right->root ---> root->right->left.
    at the end reverse result agaain to take correct result.
    But, this is not exactly postorder. This is tricky part. So we need to use 2nd option.
2. to do so we need to keep track 2 stack.
    add node to 1 stack.
    if node.left is not null, then add to stack1.
    if node.right is not null, then add to stack2.

      1
     / \
    2    3
   / \  / \
  4   5    6
          /
          2

 until s1 not empty:
  s1 = add node first.
    pop node from s1.
  if has left and right child: add them to stack1.
  s2 = add popoed node, root to stack2.
  s1 = 4,5. end
  s2 = 1,3,2,5,4
  at the end pop from stack2 and add to result, go get postoreder traversal.


      1
     / \
    2    3
   / \  / \
  4   5    6
          /
          2

  s1 = 4,5
  s2 = 1,3,6,2,2,5,4
    4,5,2,2,6,3,1

*/

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var tmps = make([]*TreeNode, 0)
	var nodes = make([]*TreeNode, 0)
	var res = make([]int, 0)

	tmps = append(tmps, root)

	for len(tmps) > 0 {
		node := tmps[len(tmps)-1]
		tmps = tmps[:len(tmps)-1]

		nodes = append(nodes, node)
		if node.Left != nil {
			tmps = append(tmps, node.Left)
		}
		if node.Right != nil {
			tmps = append(tmps, node.Right)
		}
	}

	for i := len(nodes) - 1; i >= 0; i-- {
		res = append(res, nodes[i].Val)
	}
	return res
}
