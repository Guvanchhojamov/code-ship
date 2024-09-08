package main

import "fmt"

func main(){
    prefixSums([]int{1,2,3,4,5})
}

func prefixSums(nums []int ){
    var pSum = make([]int, 0, len(nums))
    pSum = append(pSum,nums[0])
    for i:=1; i<len(nums); i++{
        pSum = append(pSum,pSum[i-1]+nums[i])
    }
    fmt.Println(pSum)
}   