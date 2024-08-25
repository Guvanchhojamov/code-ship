package main 

import "fmt"
func  main(){
    missingNumber([]int{3,0,1})
}

/*
  3 0 1 
  0 1 2 
*/
func missingNumber(nums []int) int {
    i := 0;
    j := 0;
    for  j < len(nums) {
        if nums[i] != i {
            if nums[i] >= len(nums) {
              nums[len(nums)-1], nums[i] = nums[i], nums[len(nums)-1]
            }
            fmt.Println(nums, nums[i], i )
            nums[nums[i]],nums[i] = nums[i], nums[nums[i]]
        }else{
            i++
        }
        j++
    }
    fmt.Println(nums)
    return 0
}