package main 

import "fmt"
func  main(){
   fmt.Println( missingNumber([]int{2,0,3}))
}

/*
  3 0 1 
  0 1 2 
*/
func missingNumber(nums []int) int {
    // solved with using cyrcle sort.
    i:=0 
    for i < len(nums) {
        if nums[i] < len(nums) && nums[i] != nums[nums[i]] {
            nums[i],nums[nums[i]] = nums[nums[i]],nums[i]
        }else{
            i++
        }
    }
    fmt.Println(nums)
    for i,val:=range nums {
        if i != val {
            return i 
        }
    }
    return len(nums) 
}
