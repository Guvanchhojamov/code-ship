package main

import "fmt"

func main() {
	//matrix := [][]int{
	//	{3, 7, 8},
	//	{9, 11, 13},
	//	{15, 16, 17},
	//}
	matrix2 := [][]int{
		{1, 10, 4, 2},
		{9, 3, 8, 7},
		{15, 16, 17, 12},
	}
	findLuckyNumbers(matrix2)
}

func findLuckyNumbers(matrix [][]int) {
	//rows := []int{}
	rows := []int{}
	for i := 0; i < len(matrix); i++ {
		//maximum := 0
		minimum := 1000001
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < minimum {
				minimum = matrix[i][j]
			}
		}
		rows = append(rows, minimum)
	}
	for _, row := range matrix {
		for j := 0; j < len(row); j++ {
			if j < len(row) {
				fmt.Println("cols", row[j])
			}
		}

	}

	fmt.Println(rows)
}
