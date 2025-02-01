package main

func main() {
	spiralOrder([][]int{{2, 5, 8, -1, 0, 4}})
}

func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix)-1, len(matrix[0])-1
	left, right := 0, m
	top, bottom := 0, n
	res := []int{}

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for j := top; j <= bottom; j++ {
			res = append(res, matrix[j][right])
		}
		right--
		if top <= bottom {
			for k := right; k >= left; k-- {
				res = append(res, matrix[bottom][k])
			}
			bottom--
		}
		if left <= right {
			for l := bottom; l >= top; l-- {
				res = append(res, matrix[l][left])
			}
			left++
		}
	}
	return res
}

/*
  We need here 4 loops for print nums on spiral mine
  left:=0
  right:=n-1
  top:=0
  bottom:=m-1
  1. From left to right
   for i to right:
    print num[left][i]
  2. From top to bottom
    for i:=top to bottom:
     print [i][right]  by one
     top++
  3. from i:=rigth to left
       print [top][i]
  4. from bottom to top.
    for i:=bottom to top:
      print [left][i]
    bottom--
    left++
    rigth--
  The key point is, left<=right
    Because for spiral, sprial ends in left==right position.
    While left<=right we run 4 loops like:

*/
/*
 1 2  3  4
 5 6  7  8
 9 10 11 12
 left=1
 right=1
 top=2
 bottom=1

*/
/*
 1 2 3
 4 5 6
 7 8 9
*/
/*
 6 9 7
 top=1
 left=0
 right=2
 bottom=0

*/
