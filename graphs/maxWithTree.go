package graphs

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * Definition for a binary tree node.

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Queu struct {
	node *TreeNode
	idx  int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q = []Queu{}
	var maxW = 0
	q = append(q, Queu{root, 1})

	for len(q) > 0 {
		x := len(q)
		first := q[0]
		last := q[x-1]
		for x > 0 {
			x--
			n := q[0]
			q = q[1:]
			with := last.idx - first.idx + 1
			maxW = max(maxW, with)

			// normolize index, to avoid index overflow.
			idx := n.idx - first.idx

			if n.node.Left != nil {
				q = append(q, Queu{n.node.Left, 2 * idx})
			}
			if n.node.Right != nil {
				q = append(q, Queu{n.node.Right, 2*idx + 1})
			}
		}

	}
	return maxW
}

/*
 We need to find max with of binary tree.
    what is max with.
 this is diff beetween leftmost and rightmost nodes.
    where both are not null.
    but beetween them if null this is assumed that we have node.
so with = leftmost(notnull) - rightMost(notnull) value.
    and we need to return max, so for this we needto go max down level
    as possible.
    what traversal we need to use?
 since we need leftmost and rightmost not null node in eahc level.
   we use BFS levelorder traversal.
   assume each level is 1d array:
    [1,2,4,0,6]
     0 1 2 3 4
    how we find leght? len == with.
    is included start and end? Yes.
    so if we know start, end index then len = end-start+1
    with = j-i+1. 4-0+1 =5. our length is 5.
   so to use this we need index for each node.
    so how do we know how we can set index for each node ?
  assueme start index i = 1;
  from trees whe know left child index and rightchild index.
    if start index = 0:
        leftchild = 2*i+1
        rightchild = 2*i+2
        0,1,2; 3,4; ...
    we need 1 indexed so:
        leftchild = 2*i
        rightchild = 2*i+1
      1 2; 3, 4; 5,6; ...
Ho we give index?
    since we need null nodes in between.
 we add to q node with new index.
    and we poop first take index.
    and we pop last take index.
    and if both are not null set
       maxWith =  l.i - f.i +1
    if one of them is null then we dont need calculate this level.
But, is it can index overflow?
     sice res = a-b is guaranteed in 32-bit.
     but a,b can be large..
i think no, left try implment this.

*/
