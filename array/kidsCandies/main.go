package main

import "fmt"

func main() {
	kidsCandies([]int{12, 1, 12}, 10)
}

func kidsCandies(candies []int, extraCandies int) []bool {
	max := 0
	var result = make([]bool, 0, 0)
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	fmt.Println(result)
	return result
}
