package daily
/*
You are given the two integers, n and m and two integer arrays, hBars and vBars.
The grid has n + 2 horizontal and m + 2 vertical bars,
creating 1 x 1 unit cells. 
The bars are indexed starting from 1.

You can remove some of the bars in hBars from horizontal bars and
some of the bars in vBars from vertical bars.
Note that other bars are fixed and cannot be removed.

Return an integer denoting the maximum area of a square-shaped hole in the grid,
after removing some bars (possibly none).

Input: n = 2, m = 1, hBars = [2,3], vBars = [2]
Output: 4

Explanation:
The left image shows the initial grid formed by the bars.
The horizontal bars are [1,2,3,4], and the vertical bars are [1,2,3].
One way to get the maximum square-shaped hole is by removing horizontal bar 2
and vertical bar 2.


Constraints: 
	range n m ? 
	all vbars and hbars are unique. 

*/