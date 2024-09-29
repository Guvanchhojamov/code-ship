package main

import "fmt"

func main(){
//    prefixSums([]int{1,2,3,4,5},2,3)
    subArraySumK([]int{3, 4, 7, 2, -3, 1, 4, 2}, 7)
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


/*
  nums = 3 4 7 2 -3 1 4 2    k =  7    out=4; 
 
  sum[j]-sum[i] = k  
  sum[j]-k = sum[i]

*/

func subArraySumK(nums []int, k int) int{
    pSums := map[int]int{}
    var currSum = 0
    for i:=0;i<len(nums);i++{
        currSum = currSum + nums[i]
        pSums[i] = currSum
    }

    fmt.Println(pSums, currSum)
    return 0 
}