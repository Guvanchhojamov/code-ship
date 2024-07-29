package main

func main() {
	sortArray([]int{-5, 2, 3, 1})
}

// leetcode 912.
func sortArray(nums []int) []int {
	result := []int{}
	minimum, maximum := findMinMax(nums)
	var numCounts = make([]int, maximum+1-minimum)

	for _, v := range nums {
		numCounts[v-minimum]++
	}

	for i, val := range numCounts {
		for val > 0 {
			result = append(result, i+minimum)
			val--
		}
	}
	return result
}

func findMinMax(nums []int) (int, int) {
	maximum := 0
	minimum := nums[0]
	for _, v := range nums {
		if v > maximum {
			maximum = v
		}
		if v < minimum {
			minimum = v
		}
	}
	return minimum, maximum
}
