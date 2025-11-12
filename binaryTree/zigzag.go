package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
 lets first write levelorder traversal.
    to do this we use queue ds, with is first in first out.
    we add nodes, level by level, and pop exactly x nodes, for this level.
    and add node values to level list, and add childrens to q.
*/

type Queue []*TreeNode

func (q *Queue) PopFront() *TreeNode {
	if len(*q) == 0 {
		return nil
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *Queue) Pushback(node *TreeNode) {
	*q = append(*q, node)
}
func (q *Queue) Len() int {
	return len(*q)
}
func zigzagLevelOrder(root *TreeNode) [][]int {
	var q = make(Queue, 0, 0)
	var res = make([][]int, 0, 0)
	var level = 0
	q.Pushback(root)
	for q.Len() > 0 {
		levelLen := q.Len()
		levelVals := []int{}
		for levelLen > 0 {
			node := q.PopFront()
			if node != nil {
				levelVals = append(levelVals, node.Val)
				q.Pushback(node.Left)
				q.Pushback(node.Right)
			}
			levelLen--
		}
		if len(levelVals) > 0 {
			if level%2 != 0 {
				levelVals = reverse(levelVals)
			}
			res = append(res, levelVals)
		}
		level += 1
	}
	return res
}

func reverse(nums []int) []int {
	var i, j = 0, len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

/*
ok, now we need zigzak order response, how we can do this?
- if we srart from left, we need start from right in next level. what does it means?
- take level 1. - left.
    1 - left; 2 - right; 3- left; 4-right... you see pattern?
we need level numer, we we has level
    if level % 2 == 1 then - left else from right.

*/
