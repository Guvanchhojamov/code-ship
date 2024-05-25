package main

/*
l r r
0 0 1 0 3 12

	l     r

1 0 0 0 3 12

	l      r

1 3 0 0 0 12

	l    r

1 3 12 0 0 0

end

	l r

1 2 0 4 5 6
*/

func main() {
	moveZeroes([]int{1, 2, 0, 4, 5})
}

func moveZeroes(nums []int) {

	for l, r := 0, 0; r < len(nums) && l < len(nums)-1; r++ {
		if nums[l] == 0 && nums[r] == 0 {

		}
	}
}
