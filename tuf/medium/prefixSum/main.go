package main

import "fmt"

func main(){
    prefixSums([]int{1,2,3,4,5},2,3)
}

func prefixSums(nums []int, left, right int ){
    var pSum = make([]int, 0, len(nums))
    var subSum = 0
    pSum = append(pSum,nums[0])
    for i:=1; i<len(nums); i++{
        pSum = append(pSum,pSum[i-1]+nums[i])
    }

    if left<0 || right>=len(nums) {
        fmt.Println("invalid left or right")
        return 
    }

    if left == 0{
        subSum = pSum[right]
    }else if right<len(nums){
        subSum=pSum[right]-pSum[left-1]
    }

    fmt.Println(pSum, subSum)
}   

/*
  1 2 3 4 5  k=5 , out=2 
  1 3 6 10 15 
  0:1 
  currSum +=nums[i]



*/