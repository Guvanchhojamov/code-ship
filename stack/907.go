package stack

/*
907. Sum of Subarray Minimums
Companies
Given an array of integers arr,
find the sum of min(b), where b ranges over every (contiguous) subarray of arr.
Since the answer may be large, return the answer modulo 10^9 + 7.
Example 1:

Input: arr = [3,1,2,4]
Output: 17
Explanation:
Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].
Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
Sum is 17.

Example 2:
Input: arr = [11,81,94,43,3]
Output: 444

Constraints:
1 <= arr.length <= 3 * 10^4
1 <= arr[i] <= 3 * 10^4
*/
/*
  The brute force approach is using two loops and generate each subarray and check.
tc: O(n*n)
sc: O(1)
*/

/*
Method 2: Monotic stack

	How many times a number will appear in the results(Not thinking about minimum)
	- for a element A[i], it will represents (i + 1) * (len - i + 1) times.
	- i + 1 means the subarrays whose end point is A[i]
	- (len - i + 1) means the subarrays whose start point is A[i].

How to solve this question:
  - res = sum(A[i] * f[i], A[i] is the number and f[i] means the number of subarrays where A[i] is the minimum in the subarray.
  - f[i] = left[i] * right[i], where left[i] is the left subarray numbers where A[i] is the last element and
    A[i] is minimum in current array and right[i] is the minimum and first element in the subarray.
  - How to get left[i] and right[i]. Consider leetcode question 901.
    We need to consider a corner case, two elements have the same value.

res = a[i] * left[i]*right[i]

arr = [3,1,2,4] Output: 17
Explanation:
Subarrays are:
[3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].

the a[i] contains on some subarrays, we need to focus for each a[i].

	The core idea is:

How many subarrays exist where arr[i] is the minimum?
Find the count of these subarrays and arr[i] * count. => sum+=arr[i] * count.

But the problem is how to find how many subarrays exits where arr[i] is minimum.
We need, leftMax[], and rightMax[]
leftMax = count of elements witch > a[i] on the left of a[i] including a[i].
rightMax = count of elements witch >= a[i] on the right of a[i] including a[i],
Then a[i]will be the min element on left[i] * right[i] subarrays.
[3,1,2,4] for 2:
left[2]=1; right[2]=1 = 2*1 = 2.
for 1:
left[1]=2(include 1); right[1]=3(include 1).

	a[i] * l[i]* r[i] = 2 * (1*3) = 6 - subarrays had witch min element =1.

1. Find prevMax[i] elements for i(inclusive i) count
2. Find nextMax[i] elements for i(inclusive i) count
4. iterate over array and sum prevMax and nextMax counts.
*/

func sumSubarrayMins(arr []int) int {
	var mod = 1_000_000_007
	var total = 0
	prevSmallIdx := prevSmallEqidx(arr)
	nextSmallIdx := nextSmallidx(arr)
	for i := 0; i < len(arr); i++ {
		left := i - prevSmallIdx[i]
		right := nextSmallIdx[i] - i
		subsWithMini := (arr[i] * left % mod) * right % mod
		total = (total + subsWithMini) % mod
	}
	return total

}

func nextSmallidx(nums []int) []int {
	var nsi = make([]int, len(nums), len(nums))
	var st = make([]int, 0, len(nums))
	var n = len(nums)
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] <= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			nsi[i] = n
		} else {
			nsi[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return nsi
}

func prevSmallEqidx(nums []int) []int {
	var n = len(nums)
	var psi = make([]int, n, n)
	var st = make([]int, 0, n)

	for i := 0; i < n; i++ {
		for len(st) > 0 && nums[i] < nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			psi[i] = -1
		} else {
			psi[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return psi
}
