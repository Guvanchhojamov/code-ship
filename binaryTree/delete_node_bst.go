package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		// we found deletion node
		// in this cases if one of childs are nil jsut return other, and all good.
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// if both childs not null. Hardest case all we need todo.
			/*
			   1. get right stree min val. // or get left Max node.
			   2. replace right/left min val with curr.Val
			   3. delete right/left min val.
			*/
			leftMaxNode := getLeftSubtreeMax(root.Left)
			root.Val = leftMaxNode.Val
			root.Left = deleteNode(root.Left, leftMaxNode.Val)
		}

	}
	return root
}

func getLeftSubtreeMax(node *TreeNode) *TreeNode {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

/*
 since it iw bst we go left or Right, comparing key and root.Val.
    if val < key:
        node.Left = call recursion and go left.
    elif val > key:
        node.Right = call recursion and go Right
    else : // we found key == val.  delete node.
        if node.Left == nil:
        return node.Right
        elif node.Right == nil:
        return node.Left as new val.
        as new prev right val.
    in other case  this we need to go right minimum val go connect del node   left -> right.Min.
     because it is bst.
        all left of del. node must be minimum from all right subtree.
    then we connect
        deletedNode.val =  rightMinNode.
        then delete right minNode.
    and return.
*/
