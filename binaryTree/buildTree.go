package binarytree

import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTreePreIn(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rootIndx := slices.Index(inorder, rootVal)
	if rootIndx == -1 {
		return nil
	}

	node := &TreeNode{Val: rootVal}
	node.Left = buildTree(preorder[1:rootIndx+1], inorder[:rootIndx])
	node.Right = buildTree(preorder[rootIndx+1:], inorder[rootIndx+1:])
	return node
}

/*
 preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
*/

/*
 we are given preorder and inorder traversal.
    we need build correct binary tree and return usign them.

 Preorder 1-th element always root. So we start with this.
 In inorder traversal root is always in. c r c.
  To build correctly we need, take from preorder. - its root.
  - find it from inorder.
  - left part - do again same. Inorder[prorder[i]]- root.  right part do again the same.
    appraoch:
        - iterate over preorder and take val.
        - find index of this val in inorder. It means find root.
        - take left 0->inx - as new inorder.
        - take right inx+1 -> N as another inorder right. ??
        - build new node on each step.
        - and connect left and right parts.
        - return new root node.
    - to fast look ups we create inorder map list. for finde index. key-node; val= index.
    TC: N - for map, N For iterate over preorder = 2N = N
    sc: N - for map. H- for recursive stack. N.
    lets try implement/.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(inorder []int, postorder []int) *TreeNode {
	postLastIdx := len(postorder) - 1 // shared index across recursion
	inMap := map[int]int{}
	for i, v := range inorder {
		inMap[v] = i
	}

	var helper func(inLeft, inRight int) *TreeNode
	helper = func(inLeft, inRight int) *TreeNode {
		if inLeft > inRight {
			return nil
		}

		rootVal := postorder[postLastIdx]
		root := &TreeNode{Val: rootVal}
		idx := inMap[rootVal]
		postLastIdx--

		root.Right = helper(idx+1, inRight)
		root.Left = helper(inLeft, idx-1)

		return root
	}

	return helper(0, len(inorder)-1)
}

/*
 Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]

 need to build binary tree using postorder and inorder traversal,
 - in inorder root val in the middle of each subtree.
 - in postorder the root value always at the end of each subtree.
  We can use this 2 options and take root left and right from two arrays.
  Since we know last element of postorder is root.
  appraoch is:
    - take last. of postorder as root, create new BT.
    - take rootIndx from inorder
    - in next recursion:
        since, after most right root, we have always right root, or right subtree if it is not nill
        because of this we iterate right first, then we iterate left.
    and on each iteration recursive call would be like this:
        right: [idx+1:], [:n-1] - just pop last value.
        left: [:idx], [:n-1] - just pooped last value.
            sincee we go right always first, it make sence each time we pop from last, right subtree root.
    TC: N*N - because we need to find index each time, or we can create inorder map once and use this.
        for fast lookups. Then it would be 2N
    SC: N - for map if we use, N for recursive call..
    lets try implement then. optimized version.

*/
