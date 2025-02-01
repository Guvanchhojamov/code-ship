package main

import "fmt"

func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

func maxSubArray(nums []int) int {
	mx := nums[0]
	sum := 0
	startEl := 0
	endEl := 0
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > mx {
			mx = sum
			endEl = i + 1
		}
		if sum < 0 {
			sum = 0
			startEl = i + 1
		}
	}
	fmt.Println(startEl, endEl)
	fmt.Println(nums[startEl:endEl])
	return mx
}

/*
  [-2,1,-3,4,-1,2,1,-5,4]
  mmaxim = nums[0]
  for i=0 to n:
	for j:=i to n:
	 sum+=sum[j]
	  if sum > max {
		 maximum = sum
	  }
 return maximum
*/

/*
	 mx:=nums[0]
	sum := 0
	for i:=0;i<len(nums);i++{
	 sum = 0
	 for j:=i; j<len(nums);j++{
		 sum+=nums[j]
		 if sum > mx {
			 mx=sum
		 }
	  }
	}
	return mx
*/
