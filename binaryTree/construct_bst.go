package binarytree

import (
	"slices"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	var inorder = make([]int, len(preorder))

	copy(inorder, preorder)
	sort.Ints(inorder)
	// fmt.Println(inorder)
	root := buildBST(preorder, inorder)
	return root
}

func buildBST(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	idx := slices.Index(inorder, rootVal)
	newNode := &TreeNode{
		Val: rootVal,
	}
	newNode.Left = buildBST(preorder[1:], inorder[:idx])
	newNode.Right = buildBST(preorder[idx+1:], inorder[idx+1:])
	return newNode
}

/*
 create bst from preorder traversal.
    [8,5,1,7,10,12]
 in preorder the root always first, then left then right.
    - to be bst:
        left < root < right - must be.
First:
    - sort array in asc order to take inorder traversal.
    - after use preorder and inorder to create BST.
    [8,5,1,7,10,12]
    [1,5,7,8,10,12]
- in preorder first node is root always:
    - take preorder[0]
    - take index of this idx = intorder[preorder[0]]
    - all from this inorder[:idx] - go to left.
    - all from this inorder[idx+1:] - go to right.
    - each time popfront from preorder[1:]
  newNode =inorder[:idx]
  newNode.Left =  inorder[:idx]
  newNode.Right = inorder[idx+1:]
  and continue it for each preorder value.
TC: NlogN + N
SC: N+N+N - for arrays..

*/
