package daily

/*
3160. Find the Number of Distinct Colors Among the Balls.
You are given an integer limit and a 2D array queries of size n x 2.

There are limit + 1 balls with distinct labels in the range [0, limit].
Initially, all balls are uncolored.
For every query in queries that is of the form [x, y], you mark ball x with the color y.
After each query, you need to find the number of distinct colors among the balls.

Return an array result of length n, where result[i] denotes the number of distinct colors after ith query.
Note that when answering a query, lack of a color will not be considered as a color.



Example 1:
Input: limit = 4, queries = [[1,4],[2,5],[1,3],[3,4]]
Output: [1,2,2,3]
Explanation:
After query 0, ball 1 has color 4.
After query 1, ball 1 has color 4, and ball 2 has color 5.
After query 2, ball 1 has color 3, and ball 2 has color 5.
After query 3, ball 1 has color 3, ball 2 has color 5, and ball 3 has color 4.
*/

func main() {

}

/*

  we need to store PaintQueryMap
   key-query[0]
   val-count of query frequency.

    logic:
      distinctCount = 0
	for i,val:=range queries:
      if val, exists := PaintQueryMap[query[i][0]]; !exists {
              distinctCount++
			  paintQueryMap[query[i][0]]++
       }
      res = append(res, distinctCount)

    return res

    limit = 4, queries = [[1,4],[2,5],[1,3],[3,4]]
    dc:1,2,3
    res: [1,2,2,3]
    map:
     1:3
     2:1,
     3:1

*/

func queryResults(limit int, queries [][]int) []int {
	var distinctCount int
	res := []int{}
	paintedBallColor := make(map[int]int)
	colorCount := make(map[int]int)

	for _, query := range queries {
		ball, color := query[0], query[1]
		if prevColor, exists := paintedBallColor[ball]; exists {
			if prevColor == color {
				res = append(res, distinctCount)
				continue
			}
			colorCount[prevColor]--
			if colorCount[prevColor] == 0 {
				distinctCount--
				delete(colorCount, prevColor)
			}
		}

		if colorCount[color] == 0 {
			distinctCount++
		}
		colorCount[color]++

		paintedBallColor[ball] = color
		res = append(res, distinctCount)
	}

	return res
}

/*
 [[1,4],[2,5],[1,3],[3,4]]
 dc=0
  1:0
  2:1
  2:2
  4:3
  5:4

*/

/*

  we need to store PaintQueryMap
   key-query[0]
   val-count of query frequency.

    logic:
      distinctCount = 0
	for i,val:=range queries:
      if val, exists := PaintQueryMap[query[i][0]]; !exists {
              distinctCount++
			  paintQueryMap[query[i][0]]++
       }
      res = append(res, distinctCount)

    return res

    limit = 4, queries = [[1,4],[2,5],[1,3],[3,4]]
    dc:1,2,3
    res: [1,2,2,3]
    map:
     1:3
     2:1,
     3:1

*/
