package main

func main() {

}

/*
You are given an n x n 2D matrix representing an image,
rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place,
which means you have to modify the input 2D matrix directly.
DO NOT allocate another 2D matrix and do the rotation.
1 2 3
4 5 6
7 8 9

1 4 7
2 5 8
3 6 9


1 4 7
4 5 8
7 8 9
*/
/*
 00 01 02
 10 11 12
 20 21 22
   for i:=0 to n
     fo j:=0 to m
	   nums[i][j] = swap(nums[i][j] = nums[j][i])

 for i:=0 to
*/
func rotate(matrix [][]int) {
	var n = len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		k, l := 0, n-1
		for k < l {
			matrix[i][k], matrix[i][l] = matrix[i][l], matrix[i][k]
			k++
			l--
		}
	}
}
