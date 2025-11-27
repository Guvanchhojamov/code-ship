package binarytree

import (
	"math"
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
type QNode struct {
	node *TreeNode
	row  int
	col  int
}

type NodeInfo struct {
	val, row, col int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res = make([][]int, 0)
	var q = []QNode{}
	var mp = make(map[int][]NodeInfo)
	q = append(q, QNode{root, 0, 0})
	mini, maxi := math.MaxInt, math.MinInt
	for len(q) > 0 {
		// fmt.Println(q)
		item := q[0]
		q = q[1:]

		mp[item.col] = append(mp[item.col], NodeInfo{item.node.Val, item.row, item.col})
		maxi = max(maxi, item.col)
		mini = min(mini, item.col)
		if item.node.Left != nil {
			q = append(q, QNode{item.node.Left, item.row + 1, item.col - 1})
		}
		if item.node.Right != nil {
			q = append(q, QNode{item.node.Right, item.row + 1, item.col + 1})
		}
	}

	for i := mini; i <= maxi; i++ {
		nodeInfos, ok := mp[i]
		if !ok {
			continue
		}
		if len(nodeInfos) > 0 {
			res = append(res, sortByRowOrVal(nodeInfos))
		}
	}
	return res
}

func sortByRowOrVal(infos []NodeInfo) []int {
	res := []int{}
	sort.Slice(infos, func(i, j int) bool {
		if infos[i].row != infos[j].row {
			return infos[i].row < infos[j].row
		}
		return infos[i].val < infos[j].val
	})

	for _, nodeInfo := range infos {
		res = append(res, nodeInfo.val)
	}
	return res
}

/*
 gorup by col. sort by row, if same row sort by val.
- lets try add to map whekre key= col; value=[]NodeInfos.
- at the end, we do until min to max add to response 2d array.
- if not in map then skip, not break, because keys are not continuous.
use BFS with levelorder traversal.

*/
