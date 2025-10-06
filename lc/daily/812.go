package daily

import "math"

/*
812. Largest Triangle Area

Given an array of points on the X-Y plane points where points[i] = [xi, yi],
Return the area of the largest triangle that can be formed by any three different points.
Answers within 10^-5 of the actual answer will be accepted.

Example 1:

Input: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
Output: 2.00000
Explanation: The five points are shown in the above figure. The red triangle is the largest.

Example 2:
Input: points = [[1,0],[0,0],[0,1]]
Output: 0.50000


Constraints:

3 <= points.length <= 50
-50 <= xi, yi <= 50
All the given points are unique.

*/
/*
 we are given points we need to get max triangle as possible from this points.
    need to return triangle area.
- lets assume we have 3 points, example left point, irght point, and top poont.
- the formula is the (base*height)/2
- now we need to find base and heigh to get area.
- to get max area we need to get 1 minimum point and 2 max points.
- in this sutation any base and any height are max as possible.
  take exampe for 1:
    nums = [[0,0],[0,1],[1,0],[0,2],[2,0]]
    min = [0,0]
    maxx = 2,0;
    maxy = 0,2;
    now if we got any base, or any heiht we take max possible area.
    2*2/2 = 2.00000
*/
/*
 Approach:
    take 2 vars: minpoint = [2]int
    - minx = [2]int
    - miny = [2]int
    - iterate over array:
        - check for minpoint.
        - check for minx;
        - check for miny;
    area = (minpoint[0] - minx[0]) * (minpoint[0] - miny[0]) / 2.
    return area.
*/

// func largestTriangleArea(points [][]int) float64 {
// 	var minpoint = []int{math.MaxInt, math.MaxInt}
// 	var maxx, maxy = points[0], points[0]
// 	fmt.Println(minpoint, maxx, maxy)
// 	for _, point := range points {
// 		fmt.Println(point)
// 		if point[0] <= minpoint[0] && point[1] <= minpoint[1] {
// 			minpoint[0] = point[0]
// 			minpoint[1] = point[1]
// 			continue
// 		}
// 		if point[0] > maxx[0] {
// 			maxx = point
// 			continue
// 		}
// 		if point[1] > maxy[1] {
// 			maxy = point
// 			continue
// 		}
// 	}
// 	fmt.Println(minpoint, maxx, maxy)
// 	base := maxx[0] - minpoint[0]
// 	height := maxy[1] - minpoint[1]
// 	fmt.Println(base, height, float64(base*height)/float64(2))
// 	return float64(base*height) / float64(2)
// }

func largestTriangleArea(points [][]int) float64 {
	maxArea := 0.0
	n := len(points)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				x1, y1 := points[i][0], points[i][1]
				x2, y2 := points[j][0], points[j][1]
				x3, y3 := points[k][0], points[k][1]

				// shoelace formula
				area := math.Abs(float64(
					x1*(y2-y3)+
						x2*(y3-y1)+
						x3*(y1-y2),
				)) / 2.0

				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
