package main

// 2661. First Completely Painted Row or Column
func main() {
	firstCompleteIndex([]int{1, 2, 3, 4, 5}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}) // 2
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	var rowPainted int
	var colPainted int

	var matMap = make(map[int][]int, 2)

	for i, row := range mat {
		for j := 0; j < len(row); j++ {
			matMap[row[j]] = append(matMap[row[j]], []int{i, j}...)
		}
	}

	for _, num := range arr {
		rowPainted += matMap[num][0]
		colPainted += matMap[num][1]
		if rowPainted == len(mat)
	}

	return 0
}

/*
Brute force.
   1. Hachan all row painted bolya?
    0,1,0,2,0,3..x,y -> count(x) == m bolanda
   2. Hachan all col painted bolya?
    1,0,2,0,3,0..x,y -> count(y) == n bolanda
    1 ya-da 2 yagday bolanda arr index-i gaytarmaly.

 brute force
 search each arr[i] from mat
 if finded
    k++ and l++
  if k == m or l == n return i
   assume in worth case all length are same n
  TC: t* (n*m) -> n*n*n*m = n^4
    is it ok?
    No
How can we optimize?
 All elements in arr an mat are unqiue so,
Les't try to use map.
 matMap[mat[i,j]] = []{i,j}
 for i-> to arr:
    xCount += matMap[arr[i]][0]  // increment row if not then returns 0
    yCount += matMap[arr[i][1]]  // increment col
   if xCount == n or yCount== m:
      reutrn i
  TC: m*n for create map
     m*n for iterate over arr
     so, m*n+m*n -> 2(m*n) -> n^2
  is it ok, yeap for 2d array it is standar.
  Let's implement.
*/
