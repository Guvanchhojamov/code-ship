package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
func singleNumber(nums []int) int {
	numsMap := map[int]int{}
	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			numsMap[v]++
		} else {
			numsMap[v] = 1
		}
	}
	for k, v := range numsMap {
		fmt.Println(k, v)
		if v == 1 {
			return k
		}
	}
	return 0
}
