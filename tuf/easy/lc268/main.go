package main

import "fmt"

func main() {
	missingNumber([]int{3, 0, 1})
}

/*
3 0 1
0 1 2

3 0 2
0 1 2

			if n[i] != i
			   if n[i] == len(n)
		            swap(n[i],n[len(n)-1])
		       else:
		       swap(n[n[i]], n[i])
	        else:
	        i++

1 0 2
0 1 2
*/

/*
Solution with XOR logic.

	0^a = a
	a^a =  0
	a^b = binary(a^b)
*/
func missingNumber(nums []int) int {
	res := 0
	actual := 0
	for i := 0; i < len(nums); i++ {
		actual = actual ^ i
	}

	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	actual = actual ^ len(nums)
	fmt.Println(actual, res, res^actual)

	return actual ^ res
}
