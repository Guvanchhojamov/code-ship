package sliding_window

import "fmt"

/*
Intro for sliding window and two pointer playlist. from tuf.
There few patterns for where we can use sliding window or two pointer patterns:
1. Constant window.
2. Longest subarray/substring (window size) with <condition>
3. Number of subarrays where condition <condition>
	- solution comes using pattern-2, with separating condition 2 or more parts.
4.Shortest/Minimum window <condition>
- pattern-1 used in more problems. Simple keep correct window problems.
- p-2 used median number of problems.
- in 3 also we use p1 or p2 with separating condition to parts.
- p-4 is more less problems with this related tasks.
*/
/*
 Let's solve simple example questions for each pattern.
p1+p2:
-Constant window: Print all windows with len k +
-Largest substring/subarray (size) with sum<=k.
p3:
- Number of subarrays where sum=k
p4:
- Shortest/Minimum subarray with sum >= k
*/

func windows(nums []int, k int) [][]int {
	var res = make([][]int, 0, len(nums))
	for i, j := 0, 0; j < len(nums); j++ {
		if j-i+1 == k {
			res = append(res, nums[i:j+1])
			i++
		}
	}
	return res
}

/*
 nums = 1,4,5,2,3,4,7,8  k=3
 1,4,5  4,5,2
 5,2,3  2,3,4
 4,7,8

brute force:
   1. start from i; print until i+k
    tc: n*k
    sc: 1
 sliding window:
    i,j=0
    until j==k-1 j++
    else
    i++
    until j<n;
    tc: n
    sc 1;
*/

func largestSubarray(nums []int, k int) int {
	var sum = 0
	var res = 0
	for i, j := 0, 0; j < len(nums); j++ {
		sum = sum + nums[j]
		if sum >= k {
			sum = sum - nums[i]
			i++
		}
		fmt.Println(i, j, sum)
		res = max(res, j-i+1)

	}
	return res
}

/*
 nums = 2,5,1,7,10  k=14
 2,5,1,7 = 15 > k
 5,1,7 = 13 < k
 other subarrays are small length so our answer is 3.
 brute force:
 start from 0 and generete all subarrays, get sum and pick max:
 tc: n*n
 sc: 1
 can we optimize tc using sliding window?
 start from 0 sum while sum <=k else
    get length of window and decrease from front.
    left, right. right until n;
    sum with prefix sum each num
    and decrement from sum when shrink the window.
*/
