package binarytree

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// Hardcoded tree:
	//
	//            1
	//        /       \
	//       2         3
	//     /   \     /   \
	//    4     5   6     7
	//         / \
	//        8   9

	n8 := &TreeNode{Val: 8}
	n9 := &TreeNode{Val: 9}
	n4 := &TreeNode{Val: 4}
	n6 := &TreeNode{Val: 6}
	n7 := &TreeNode{Val: 7}
	n5 := &TreeNode{Val: 5, Left: n8, Right: n9}
	n2 := &TreeNode{Val: 2, Left: n4, Right: n5}
	n3 := &TreeNode{Val: 3, Left: n6, Right: n7}
	root := &TreeNode{Val: 1, Left: n2, Right: n3}

	res := boundaryTraversal(root)
	fmt.Println("Boundary traversal:", res)
}

/*
Given the root of a binary tree, Find its boundary traversal in anti-clockwise order, starting from the root.

The boundary of a binary tree consists of the following parts:

Left Boundary: The nodes on the left edge of the tree, excluding the leaf nodes.
Leaf Nodes: All the leaf nodes from left to right.
Right Boundary: The nodes on the right edge of the tree, excluding the leaf nodes, traversed in bottom-up order.
Note:

If the root does not have a left subtree or right subtree, then the root itself is considered part of the respective boundary.
Each node in the boundary should appear only once in the traversal.
*/
/*
 ok, to get bundary traversal, we need:
    1. left bundary vals.
    2. leaf node vals.
    3. right bundary node vals.

*/
func boundaryTraversal(root *TreeNode) []int {
	// 1. collect left bundary exepct leaf nodes.
	// 2. collect leaf node vals.
	// 3. collect right bundary.
	if root == nil {
		return []int{}
	}
	var res = make([]int, 0)
	var left = make([]int, 0)
	var right = make([]int, 0)
	var leaf = make([]int, 0)

	res = append(res, root.Val)

	collectLeft(root.Left, &left)

	collectLeaf(root.Left, &leaf)
	collectLeaf(root.Right, &leaf)

	collectRight(root.Right, &right)

	res = append(res, left...)
	res = append(res, leaf...)
	res = append(res, right...)
	return res
}

/*
collecting of the left bundary we do with preorder traversal.
prioritizing go left, always, if left is nil then go right.
if it is leaf node, then return .
*/
func collectLeft(node *TreeNode, left *[]int) {
	if node == nil {
		return
	}
	// check is it leaf node.
	if node.Left == nil && node.Right == nil {
		return
	}

	*left = append(*left, node.Val)

	if node.Left == nil {
		collectLeft(node.Right, left)
	} else {
		collectLeft(node.Left, left)
	}
	return
}

// to collect leaf nodes, we use inorder traversal, adding only children to result array.
func collectLeaf(node *TreeNode, leaf *[]int) {
	if node == nil {
		return
	}

	// use inorder traversal, but add only if it is leaf node.
	collectLeaf(node.Left, leaf)

	if node.Left == nil && node.Right == nil {
		*leaf = append(*leaf, node.Val)
	}

	collectLeaf(node.Right, leaf)
	return
}

// to collect right bundary we go in postorder traversal, to get reversed result from bottom to top.
func collectRight(node *TreeNode, right *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		return
	}

	if node.Right != nil {
		collectRight(node.Right, right)
	} else {
		collectRight(node.Left, right)
	}
	*right = append(*right, node.Val)
	return
}

/*
 to collect right bundary we go right,
 prioriitizing right bundary, and exept leaf nodes.
*/
