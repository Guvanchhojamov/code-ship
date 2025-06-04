package stack

/*

2104. Sum of Subarray Ranges

You are given an integer array nums.
The range of a subarray of nums is the difference between the largest and smallest element in the subarray.
Return the sum of all subarray ranges of nums.
A subarray is a contiguous non-empty sequence of elements within an array.

Example 1:

Input: nums = [1,2,3]
Output: 4
Explanation: The 6 subarrays of nums are the following:
[1], range = largest - smallest = 1 - 1 = 0
[2], range = 2 - 2 = 0
[3], range = 3 - 3 = 0
[1,2], range = 2 - 1 = 1
[2,3], range = 3 - 2 = 1
[1,2,3], range = 3 - 1 = 2
So the sum of all ranges is 0 + 0 + 0 + 1 + 1 + 2 = 4.
Example 2:
Input: nums = [1,3,3]
Output: 4
Explanation: The 6 subarrays of nums are the following:
[1], range = largest - smallest = 1 - 1 = 0
[3], range = 3 - 3 = 0
[3], range = 3 - 3 = 0
[1,3], range = 3 - 1 = 2
[3,3], range = 3 - 3 = 0
[1,3,3], range = 3 - 1 = 2
So the sum of all ranges is 0 + 0 + 0 + 2 + 0 + 2 = 4.

Example 3:
Input: nums = [4,-2,-3,4,1]
Output: 59
Explanation: The sum of all subarray ranges of nums is 59.

Constraints:

1 <= nums.length <= 1000
-10^9 <= nums[i] <= 10^9
*/

/*
 Need to find range for subarray. and sum of the ranges.
- we need min and max element of subarray.
nums = [1,2,3]; out=4
Explanation: The 6 subarrays of nums are the following:
[1], range = largest - smallest = 1 - 1 = 0
[2], range = 2 - 2 = 0
[3], range = 3 - 3 = 0
[1,2], range = 2 - 1 = 1
[2,3], range = 3 - 2 = 1
[1,2,3], range = 3 - 1 = 2
So the sum of all ranges is 0 + 0 + 0 + 1 + 1 + 2 = 4.
What is brute force?
- in two loops generate all subarrays:
- find min and max when generating
- subtract max-min; add to sum; return sum.
tc: O(n*n)
sc: (1)
Follow-up: Could you find a solution with O(n) time complexity?
left smallest element
right smallest element

*/
/*

For each element nums[i], determine:

How far you can go left (Previous Greater/Lesser).

How far you can go right (Next Greater/Lesser).

The number of subarrays where nums[i] is the max/min is:

count = (i - left) * (right - i)
contribution = count * nums[i]

Final Result:
Add up contributions where nums[i] is max.
Subtract contributions where nums[i] is min.
*/
// Kakoy schitaetsya team lead, kotoryy pozval druga v komandu razrabotchikov kotoryy  7 let rabotal v gos,
// bez intervyu, bez sprosa u drugih, i drigie dolzhny ego uchit, zastavit sebya chtoby
// s nim rabotat, i tyanut, a sam chillit, delate chto hochet. A etot drug za 6 mezyazev tolkom nichego ne uchilsya tolko
// delaet vid. Chto delat, kakoy vyvod sdelat?

/*
  instead of generating all subarrays, we need compute:
- how many subarrays with min is a[i]
- how many subarrays with max is a[i]
  to count this:
for each  i:
countMin = (i-leftMin[i])*(rightMin[i]-i)  (we assumue leftMin and rightMin contains indesex array)
countMax = (i-leftMax[i])*(rightMax[i]-i)  (we assume helper precomputed arrays)
maxContributions = countMax * arr[i]
minContributions = countMin * a[i]
total = sum of all maxContrib - sumOff all MinContributions.
or we can do this:
total+=maxContr  - add up
total-=minContr  - and subtract for each arr[i]
*/

func subArrayRanges(nums []int) int64 {
	//var largestsSum int64 = 0
	//var smallestsSum int64 = 0
	var total int64 = 0
	leftMins := prevMinIdxs(nums)
	rightMins := nextMinIdxs(nums)
	leftMaxs := prevMaxIdxs(nums)
	rightMaxs := nextMaxIdxs(nums)

	for i := 0; i < len(nums); i++ {
		contributionsMin := ((i - leftMins[i]) * (rightMins[i] - i)) * nums[i]
		contributionsMax := ((i - leftMaxs[i]) * (rightMaxs[i] - i)) * nums[i]
		//largestsSum += int64(contributionsMax)
		//smallestsSum += int64(contributionsMin)
		total += int64(contributionsMax)
		total -= int64(contributionsMin)
	}
	return total
}

func prevMinIdxs(nums []int) []int {
	var prevMins = make([]int, len(nums), len(nums))
	var st = make([]int, 0, len(nums))
	var n = len(nums)
	for i := 0; i < n; i++ {
		for len(st) > 0 && nums[i] < nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			prevMins[i] = -1
		} else {
			prevMins[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return prevMins
}

func nextMinIdxs(nums []int) []int {
	var n = len(nums)
	var nextMins = make([]int, n, n)
	var st = make([]int, 0, n)

	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] <= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			nextMins[i] = n
		} else {
			nextMins[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return nextMins
}

func prevMaxIdxs(nums []int) []int {
	var n = len(nums)
	var prevMax = make([]int, n, n)
	var st = make([]int, 0, n)

	for i := 0; i < n; i++ {
		for len(st) > 0 && nums[i] > nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			prevMax[i] = -1
		} else {
			prevMax[i] = st[len(st)-1]
		}
		st = append(st, i)
	}

	return prevMax
}

func nextMaxIdxs(nums []int) []int {
	var n = len(nums)
	var nextMax = make([]int, n, n)
	var st = make([]int, 0, n)

	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] >= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			nextMax[i] = n
		} else {
			nextMax[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return nextMax
}
