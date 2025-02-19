package binarysearch

/*

Given an array/list of length ‘N’,
where the array/list represents the boards
and each element of the given array/list represents
the length of each board.
Some ‘K’ numbers of painters are available to paint these boards.
Consider that each unit of a board takes 1 unit of time to paint.
You are supposed to return the area of the minimum time
to get this job done of painting all the ‘N’ boards
under the constraint that any painter will only paint the continuous sections of boards.

Example 1:
boards[] = {5, 5, 5, 5}, k = 2
Result:
 10
Explanation:
 We can divide the boards into 2 equal-sized partitions,
 so each painter gets 10 units of the board and the total time taken is 10.

 Example 2:
 boards[] = {10, 20, 30, 40}, k = 2
Result: 60
Explanation:
 We can divide the first 3 boards for one painter and the last board for the second painter.
*/

/*
Divide board to k painters such that:
 - Need to paint all boards
 - And need to taken minimum time for paint all
 - need divide boards in continous manner.

 What is the maximum time to paint all board?
	For example we have 1 pointer:
	We need to be give all boards to 1 painter so,
	sum(arr) is the max time or worth case.
 What is the minimum time can take paiting?
	For example we have 5 boards and 5 pointers:
	Then it is the maximum board timm, so
	max(arr) it is the best case.
	So our range is
	from max(arr) ... to sum(arr)
*/

func paintersPartition(nums []int, k int) int {

	a := max(nums...)

}
