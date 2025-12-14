package graphs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/// more readable version
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	return dfs(root, p, q)
// }

// func dfs(node, p, q *TreeNode) *TreeNode {
// 	if node == nil {
// 		return nil
// 	}

// 	if node == p || node == q {
// 		return node
// 	}

// 	left := dfs(node.Left, p, q)
// 	right := dfs(node.Right, p, q)

// 	if left != nil && right != nil {
// 		return node
// 	}

// 	if left == nil {
// 		return right
// 	}
// 	return left
// }

/*
ok we need to find common arcester for p,q nodes.
 - node itself are ancestor.
 - all nodes are unique.
 - p,q is exists in tree.
    Whan we need to do:
 Since we need to find common, the path to p and path to q,
    somehow must be similar.
    and somewhere they seperated.
    we need to find commond so we need to find seperation point
    and take node befor seperation, this is our lowest ancestor for p,q.
   - Since we take last common it is guaranteed it is lowest.
 How iterate ?
  - BFS or DFS ?
    since we need to go depth for each node. we dont need traverse level by level.
    we need 1 node on each path.
    so use DFS.
Example:
    p=6; q=7;
path:
 p = [3,5,6]
 q = [3,5,2,7]
 so, now we can see common nodes in each path.
    - assume we can paste -1 for no nodes.
    - take max length. and iterate.
 - if p[i] == q[i] then continue:
    where p[i] != q[i] we find seperation and return
        node with position i-1.
    TC: N for p + N for Q + N for iteration. 3N = N.
    sc: N for p + N for Q = 2N = N
How can we optimize?
    Since we search, node from left and right.
    we except something from left and right.
    where from left and right comes some value.
        this is our common node.
    so lets take example:
        1
      2    7
    3   4
     5   6
p=5; q = 6;
 - use DFS again.
 - go left. if we found p in left subtree return it
 - go ring. if we found q in right subtree return it.
  - return nil if we dont found anything.
 1->2->3-> <-nul,
        ->5 = p <-5.
   we on 3:
    left: nil
    right: 5
   if we on node where one of left or right are null return not null val.
   1->2:
   left:5=p;
   go to right.
   1->2->4-><-null
          -><-6 = q
   1->2:
        left = 5 not null
        right = 6 not null
    in that case we found min common ancestor; return curr node Val.
    as result.
    tc: N - for traversal.
    SC: N for stack.
 this is much better.
Since we need optimal i will go with this way:
    - understand both. what we did.
    - implement only optimal, because this is most expected in interviews...
    so lets implement second.
*/
