package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum2(1000))
}

/*
Input: c = 5
Output: true
Explanation: 1 * 1 + 2 * 2 = 5
*/

func judgeSquareSum(c int) bool {
	multiMap := map[int]int{}
	res := c
	if c > 10 {
		sqrtRes := math.Sqrt(float64(c))
		res = int(math.Round(sqrtRes))
	}
	for i := 0; i <= res; i++ {
		multiMap[i*i] = i
	}
	for val, _ := range multiMap {
		if _, ok := multiMap[c-val]; ok {
			return true
		}
	}
	return false
}
