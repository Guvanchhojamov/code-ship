package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minDiffLowAndHigh([]int{82, 81, 95, 75, 20}))
}

/*
time complexity: O(N+N+N) = O(N)
space complexity: O(n)

20 75 81 82 95

6 0 1 4 5 6 6

0 1 4 6 6 6

1,5,0,10,14

0 1 5 10 14
*/
func minDiffLowAndHigh(nums []int) int {

	if len(nums) < 3 {
		return 0
	}

	var tmpNums = make([]int, len(nums)) //O(n)

	slices.Sort(nums) // O(N)
	minimum, maximum := nums[0], nums[len(nums)-1]

	copy(tmpNums, nums)

	ConvToMinResult := convertLastThreeToMin(nums, minimum)     //O(N)
	ConvToMaxResult := convertFirtsThreeToMax(tmpNums, maximum) // O(N)

	return min(ConvToMinResult, ConvToMaxResult)

}

func convertLastThreeToMin(nums []int, min int) int {
	nums[len(nums)-1], nums[len(nums)-2], nums[len(nums)-3] = min, min, min
	return getMax(nums) - min
}

func convertFirtsThreeToMax(nums []int, max int) int {
	nums[0], nums[1], nums[2] = max, max, max
	return max - getMin(nums, max)
}

func getMax(nums []int) int {
	tmpMax := 0
	for _, v := range nums {
		if v > tmpMax {
			tmpMax = v
		}
	}
	return tmpMax
}

func getMin(nums []int, max int) int {
	tmpMin := max
	for _, v := range nums {
		if v < tmpMin {
			tmpMin = v
		}
	}
	return tmpMin
}
