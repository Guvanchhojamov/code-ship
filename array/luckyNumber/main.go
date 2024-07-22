package main

import "fmt"

func main() {
	// matrix := [][]int{
	// 	{3, 7, 8},
	// 	{9, 11, 13},
	// 	{15, 16, 17},
	// }
	matrix2 := [][]int{
		{1, 10, 4, 2},
		{9, 3, 8, 7},
		{15, 16, 17, 12},
	}
	findLuckyNumbers(matrix2)
}

// func findLuckyNumbers(matrix [][]int) []int {
// 	m := len(matrix)
// 	maxs := []int{}
// 	mins := []int{}
// 	//result := []int{}
// 	for i := 0; i < m; i++ {
// 		n := len(matrix[i])
// 		maximum := 0
// 		minimum := 1000000
// 		for j := 0; j < n; j++ {
// 			if matrix[i][j] < minimum {
// 				minimum = matrix[i][j]
// 			}
// 		}
// 		mins = append(mins, minimum)
// 		fmt.Println("i", i)
// 		for k := 0; k < len(matrix); k++ {
// 			if matrix[k][i] > maximum {
// 				maximum = matrix[k][i]
// 			}
// 		}
// 		maxs = append(maxs, maximum)
// 		//fmt.Println("min:", minimum, "max:", maximum)
// 	}

// 	fmt.Println("mins", mins)
// 	fmt.Println("mins", maxs)
// 	return nil
// }

func findLuckyNumbers(matrix [][]int) {
	rows := []int{}
	for i, row := range matrix {
		//maximum := 0
		minimum := 1000000
		for j, item := range row {
			if item < minimum {
				minimum = item
			}
			fmt.Println(matrix[j][i])
		}
		rows = append(rows, minimum)

		// for k := 0; k < len(row); k++ {
		// 	fmt.Println(matrix[i][k])
		// }
	}
	fmt.Println(rows)

	fmt.Println(rows)
}
