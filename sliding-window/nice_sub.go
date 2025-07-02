package sliding_window

/*
1248. Count Number of Nice Subarrays

Given an array of integers nums and an integer k. A continuous subarray is called nice if there are k odd numbers on it.

Return the number of nice sub-arrays.



Example 1:

Input: nums = [1,1,2,1,1], k = 3
Output: 2
Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].
Example 2:

Input: nums = [2,4,6], k = 1
Output: 0
Explanation: There are no odd numbers in the array.
Example 3:

Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
Output: 16


Constraints:

1 <= nums.length <= 50000
1 <= nums[i] <= 10^5
1 <= k <= nums.length
*/

func numberOfSubarrays(nums []int, k int) int {
	var l, m, r = 0, 0, 0
	var res = 0
	var odds = 0

	for r < len(nums) {
		if nums[r]%2 != 0 {
			odds++
		}

		for odds > k {
			if nums[l]%2 != 0 {
				odds--
			}
			l++
			m = l
		}

		if odds == k {
			for m < len(nums) && nums[m]%2 == 0 {
				m++
			}
			res = res + (m - l) + 1
		}
		r++
	}
	return res
}

/*
 Given need to return number of subarrays witch contains exactly k odd number.
 Bture force:
1. generate all subarrays and check odd numbers count.
	generate in 2 for loops.
	tc: n*n
	sc: O(1)
2. how can we optimze this?
	We can replace the odd nums with 1 and even nums with 0.
	The count off odd nums = the sum off odd nums in there.
	so we can where sum <=k and sum <= k-1
	and subtract x-y.
	tc: n+n
	sc: 1
3. using 3 pointer sliding window.
	left, middle, right.

[2,2,1,1,2,1,2]  k =3; res = 6.
l= 0
m= 2
r= 5
odds=3.
res=0
  res = res + m-l+1  = 2+
 until r< n:
	if odd > k:
		while odds>k:
			if a[l] is odd odds--
			l++
			m = l

	if odd == k:
		while a[m] is not odd:
			m++
	  res = m-l+1
 r++
 tc: n
 sc:n
*/
