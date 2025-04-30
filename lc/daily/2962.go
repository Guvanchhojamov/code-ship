package daily

/*
2962.
Count Subarrays Where Max Element Appears at Least K Times
You are given an integer array nums and a positive integer k.
Return the number of subarrays where the maximum element of nums appears at least k times in that subarray.

A subarray is a contiguous sequence of elements within an array.


Example 1:

Input: nums = [1,3,2,3,3], k = 2
Output: 6
Explanation: The subarrays that contain the element 3 at least 2 times are:
 [1,3,2,3], [1,3,2,3,3], [3,2,3], [3,2,3,3], [2,3,3] and [3,3].

Example 2:
Input: nums = [1,4,2,1], k = 3
Output: 0
Explanation: No subarray contains the element 4 at least 3 times.

Constraints:

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^6
1 <= k <= 10^5
*/
/*
  What we need? we need subbarays wich contains max element at least k times.
   Conditions:
  to be at least k times the max element must appear at least k or more time.
	so, if max el freq < k return 0
 Can we generate subarrays and check?
	for m:=0 to n:
      for i:m to n:
		 for j:0 to i:
			maxCount++
      if maxCount>=k res++; maxCount = 0
tc: n^3
sc: 1
 Since our range is too large we got tle in this solution.
How can we optimize?
	[1,3,2,3,3]
take thees example we have 3 max numbers.
	using this 3 max number how much subarrays can generate?
and matter where they placed.
   since topic is sliding window how can we use window in there?

	1. iterate over array anc find and count max.
       if maxCount<k return 0;
	otherwise add res+1 // because we can generate at least 1 subarray, and lets assume this whole subarray.
    2. iterate over array again and keep and count window.
		how we move and keep window, and calculate correct the maxNumsCount.
[1,3,2,3,3]
   i,j = 0,0
	for i < n :
     if a[j] == mx && j+1>=k:
 		mxCount++

     if mxCount < k && j<n:
		j++
     else:
      i++
       if a[i] != mx:
		  mxCount--
     if mxCount == k:
        res++
0,0
1
so i cant solve this right know.
*/
