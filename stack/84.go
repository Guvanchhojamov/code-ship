package stack

import "fmt"

/*
84. Largest Rectangle in Histogram
Given an array of integers heights representing the histogram's bar
height where the width of each bar is 1, return the area of the largest rectangle in the histogram.

Example 1:
Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.

Example 2:

Input: heights = [2,4]
Output: 4

Constraints:

1 <= heights.length <= 10^5
0 <= heights[i] <= 10^4
*/

/*
Input: heights = [2,1,5,6,2,3]
Output: 10

how we find the rectangle?
- to generate rectangle from heights, the height of each bar must be equal, we need cut equal heights.
- take each bar, go to left until previous min height, and go to right until next min height.
1. find previous smaller element - psei (index not the number)
2. find next smaller element - nsei (index)
3. find with for rectangle.
height = nums[i]
width  = i-psei[i]-1
area = height*weight
res = max(res,area)
for each element we need to find area and compare with max result
tc: O(3n)
sc: O(2N)
*/
func largestRectangleArea(heights []int) int {
	var largeArea int
	prevSmalls := prevSmallIdxs(heights) // take there edge case <=.
	nextSmalls := nextSmallIdxs(heights) // or there.

	fmt.Println(prevSmalls)
	fmt.Println(nextSmalls)

	for i, h := range heights {
		rectangle := h * (nextSmalls[i] - prevSmalls[i] - 1)
		largeArea = max(rectangle, largeArea)
	}
	return largeArea
}

func prevSmallIdxs(nums []int) []int {
	var prevMins = make([]int, len(nums), len(nums))
	var st = make([]int, 0, len(nums))

	for i := range nums {
		for len(st) > 0 && nums[i] <= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			prevMins[i] = -1
		} else {
			prevMins[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return prevMins
}

func nextSmallIdxs(nums []int) []int {
	var nextSmalls = make([]int, len(nums), len(nums))
	var st = make([]int, 0, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] < nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			nextSmalls[i] = len(nums)
		} else {
			nextSmalls[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return nextSmalls
}

/*
Can we optimze tc and sc makign psei and nsei in fly?
- start fron begin to end.
- keep stack monothonic in increasing order.
- if nums[i] is smaller than st.top - we found nextSmallIdx, so
- start st.pop until nums[i]<st.top calculate wigth for
- with = i-popedIdx-1; height = popedHeight; area = with*popedHeight.
- like this we caclucate are until prevSmall, and check for max. on prevSmall it will be maxArea automatically.
- is st is empty with = i - becose no prevSmall idx so we can go until 0 idx -> i rectangle.
*/
/*
 No f*cking not understandable stuff. What a f*ck.
*/
func largestRectangleAreaOptimal(heights []int) int {
	heights = append(heights, 0) // to calculate remaing area on the end,
	var st = make([]int, 0, len(heights))
	st = append(st, -1) // simplify calculation
	var maxArea = 0
	for i := range heights {
		for len(st) > 1 && heights[i] <= heights[st[len(st)-1]] {
			currHeight := heights[st[len(st)-1]]
			st = st[:len(st)-1]
			with := i - st[len(st)-1] - 1
			maxArea = max(maxArea, currHeight*with)
		}
		st = append(st, i)
	}
	return maxArea
}
