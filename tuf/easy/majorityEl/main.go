package main

import "fmt"

/*
 if sequence of element on arr bigger than n/2 then majority.
*/

func main() {
	mel := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	fmt.Println(mel)
}

/*
169. Majority Element
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

*/

/*
	Input: nums = [2,2,1,1,1,2,2]
	Output: 2
*/
func majorityElement(nums []int) int {
	numsMap := map[int]int{}
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			numsMap[num]++
		} else {
			numsMap[num] = 1
		}
	}

	fmt.Println(numsMap)
	for number, numberCount := range numsMap {
		fmt.Println(number, numberCount)
		if numberCount > len(nums)/2 {
			return number
		}
	}

	return -1
}
