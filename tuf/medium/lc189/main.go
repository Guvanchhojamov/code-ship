package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

// 189. Rotate Array
/*
	for k > 0 ...

	i=0 .. n-1

    1.swap nums[i] = [n-1] i++
*/
func rotate(nums []int, k int) {
	for k > 0 {
		for i := 0; i < len(nums); i++ {
			nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
		}
		k--
	}
	//fmt.Println(k)
}

// this solution time limit exeptions on big test case.
// need optimization
//
/*
	1 2 3 4 5 6 7   k=3    output: 5,6,7,1,2,3,4
    1 2 3 4 5 6 7 0 0 0 0
	1 2 3 4 5 6 7 1 2 3 4
	        5 6 7 1 2 3 4


     2) [-1,-100,3,99], k = 2
	     [3,99,-1,-100]

		 -1 -100 3 99 0 0
		 -1 -100 3 99 -1 -100
		         3 99 -1 -100


*/

// func rotate(nums []int, k int) {
// 	if k > 0 {
// 		for i := 0; i < len(nums); i++ {
// 			nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
// 		}
// 		rotate(nusm, k--)
// 	}
// 	return
// }

func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[(i+k)%n] = nums[i]
	}
	copy(nums, res)
}

func rotate3(nums []int, k int) {
	k = k % len(nums) // why we do this? For k always be in nums range;

	//1. Reverse all array
	reverseArray(nums)

	//2. Reverse first part
	reverseArray(nums[:k])
	//2. Reverse second part
	reverseArray(nums[k:])

}

func reverseArray(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

/*
k=2 %4 = 2

1) 1 2 3 4
2) 4 3 2 1
3) 3 4 2 1
4) 3 4 1 2

*/
