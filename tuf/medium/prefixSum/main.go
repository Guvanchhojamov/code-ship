package main

import "fmt"

func main() {
	//    prefixSums([]int{1,2,3,4,5},2,3)
	subArraySumK([]int{1, 2, 1, 2, 1}, 3)
}

func prefixSums(nums []int, left, right int) {
	var pSum = make([]int, 0, len(nums))
	var subSum = 0
	pSum = append(pSum, nums[0])
	for i := 1; i < len(nums); i++ {
		pSum = append(pSum, pSum[i-1]+nums[i])
	}

	if left < 0 || right >= len(nums) {
		fmt.Println("invalid left or right")
		return
	}

	if left == 0 {
		subSum = pSum[right]
	} else if right < len(nums) {
		subSum = pSum[right] - pSum[left-1]
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

  0 1+1   1
  7 1+1
  11 1+1   3
  13 1 + 1  4
          2


1, 2, 1, 2, 1
*/

func subArraySumK(nums []int, k int) int {
	pSusm := map[int]int{}
	count, diff, currSum := 0, 0, 0
	pSusm[0] = 1
	for i := 0; i < len(nums); i++ {
		currSum += nums[i]
		diff = currSum - k
		if val, ok := pSusm[diff]; ok {
			count += val
		}
		pSusm[currSum]++

	}
	return count
}
