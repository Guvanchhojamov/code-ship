package main

import "fmt"

func main() {
	fmt.Println([]int{1, 2, 3, 4, 5})
}

func pivotIndex(nums []int) int {
	pSums := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		pSums = append(pSums, pSums[i-1]+nums[i])
	}

	for i := range nums {
		leftSum := getSum(0, i-1, pSums)
		rightSum := getSum(i+1, len(nums)-1, pSums)
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

func getSum(l, r int, pSums []int) int {
	switch {
	case l > r:
		return 0
	case l > 0:
		return pSums[r] - pSums[l-1]
	}
	return pSums[r]
}
