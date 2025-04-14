package daily

/*
2226. Maximum Candies Allocated to K Children.

You are given a 0-indexed integer array candies.
Each element in the array denotes a pile of candies of size candies[i].
You can divide each pile into any number of sub piles,
but you cannot merge two piles together.

You are also given an integer k.
You should allocate piles of candies to K children such that each child gets the same number of candies.
Each child can be allocated candies from only one pile of candies and some piles of candies may go unused.

Return the maximum number of candies each child can get.



Example 1:
Input: candies = [5,8,6], k = 3
Output: 5
Explanation: We can divide candies[1] into 2 piles of size 5 and 3, and candies[2] into 2 piles of size 5 and 1.
We now have five piles of candies of sizes 5, 5, 3, 5, and 1.
We can allocate the 3 piles of size 5 to 3 children. It can be proven that each child cannot receive more than 5 candies.

Example 2:
Input: candies = [2,5], k = 11
Output: 0
Explanation: There are 11 children but only 7 candies in total, so it is impossible to ensure each child receives at least one candy. Thus, each child gets no candy and the answer is 0.


Constraints:
1 <= candies.length <= 10^5
1 <= candies[i] <= 10^7
1 <= k <= 10^12
*/
/*
   given candies, need divide each candy to some val, must contain one max val.
   need k max values from each basket, fo k children.
      Key points:
		- some baskets, can be unused.
        - if sum of candies < than k children, return 0, no way to divide.
      what is the min val can take each child?
		min of the candies.
        	 len(candies) == k then result= min(candies)
            15 15 25   k=3
	   what is the max resp?
        max(candies) when k=1
        45 55 75  k=1 res = 75
	So the range is
    min = min(candies).
	max = max(candies).
	result must in range min...max
		candies = [5,8,6], k = 3
       min=5, max=8; i=5,6,7,8
	    1. if sum(candies) < k return 0
		2. if len(candies) == k return min
		3. if  k==1  return  max
		otherwise:
	divide each basket for min parts.
		divide 5 15 25   k=5
        5 / 5  = 1
        15 /5  = 3
        25 / 5 = 5
       5/6 = 0
       15/6 = 2
       25 / 6 = 4

      4,7,5   k=4
      4/3 =1
      7/3 =2
      5/3 =1
	using binary search on answers.
*/

func maximumCandies(candies []int, k int64) int {
	var mx = candies[0]

	for _, val := range candies {
		mx = max(mx, val)
	}

	var subpiles int64 = 0
	var low, high = 0, mx
	for low <= high {
		mid := low + ((high - low) / 2)
		subpiles = getSubpilesCount(candies, mid)
		if subpiles >= k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}

func getSubpilesCount(nums []int, pileSize int) int64 {
	var res = 0
	for _, val := range nums {
		res += val / pileSize
	}
	return int64(res)
}
