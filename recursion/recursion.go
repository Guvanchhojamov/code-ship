package recursion

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}

/*
 We need to return indexes of the target number in array, otherwise return empty arr.
  1 3 2 2 4
  0 1 2 3 4
  res = [2,3]
*/

func targetIdx(nums []int) {

}

func fn(nums []int, i, target int) []int {
	list := []int{}
	if i == len(nums) {
		return list
	}

	if nums[i] == target {
		list = append(list, i)
	}

	ansList := fn(nums, i+1, target)
	fmt.Print(ansList)
	fmt.Println()
	list = append(list, ansList...)
	fmt.Println(list)
	return list
}

func printTriangle(n int) {
	if n == 1 {
		fmt.Println("*")
	}

}
