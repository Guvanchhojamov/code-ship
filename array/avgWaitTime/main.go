package main

import "fmt"

func main() {
	fmt.Println(averageWaitingTime(
		[][]int{
			{2, 3},
			{6, 3},
			{7, 5},
			{11, 3},
			{15, 2},
			{18, 1},
		},
	))
}

func averageWaitingTime(customers [][]int) float64 {
	orderComplitedTime := customers[0][0] + customers[0][1]
	result := customers[0][1]

	for i := 1; i < len(customers); i++ {
		arrivalTime, prepTime := customers[i][0], customers[i][1]

		if orderComplitedTime < arrivalTime {
			orderComplitedTime = arrivalTime + prepTime
			result += prepTime
			continue
		}
		orderComplitedTime += prepTime
		result += orderComplitedTime - arrivalTime
	}
	return float64(result) / float64(len(customers))
}
