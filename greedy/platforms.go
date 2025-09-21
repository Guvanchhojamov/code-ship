package greedy

import "sort"

/*
We are given two arrays that represent the arrival and departure times of trains that stop at the platform.
We need to find the minimum number of platforms needed at the railway station so that no train has to wait.

Examples 1:

Input: N=6,
arr[] = {9:00, 9:45, 9:55, 11:00, 15:00, 18:00}
dep[] = {9:20, 12:00, 11:30, 11:50, 19:00, 20:00}

Output:3
Explanation: There are at-most three trains at a time. The train at 11:00 arrived but the trains which had arrived at 9:45 and 9:55 have still not departed. So, we need at least three platforms here.

*/
/*
  - need to find min number of platforms.
  - arr and dep always equal.
  - at least we need 1 platform anyway.
1. brute force:
	How can we so this in simple way?
- we need max dep time, if next arr time is after max dep time then we dont need another platform.
- if
*/
//1. in two loops. count the overlap times. and keep track max number of overlaps. and this is the our min anser.
// how we cot overlap?
// if (arr[i] >= arr[j] && arr[i]<=dep[j]) or ((arr[j]>=arr[i]) and arr[j] <= dep[i])
func platforms(arr []int, dep []int) int {
	var count = 1
	var res = 1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(dep); j++ {
			if i == j {
				continue
			}
			if arr[j] <= dep[i] && arr[i] <= dep[j] {
				count++
			}
		}
		res = max(res, count)
	}
	return res
}

/*
  We need to use two poitner approach.
  We need to cach max trains in station at the same time.
  the max trains at station in same time == needed platforms to handle trains.
  keep track i,j pointer for i - train arrival time, j - train departure time.
  We need to sort both array independently, to hande arrival times and departure times.
  after sorting the we can see:
	- if arrival is before dep time - we need additional platform. pl+1
	- else we can free this one platform. pl-1
  The max needed platforms are our answer, because we need platforms count where
	all arrived trains handled at same time. or we can handle in same time arrived trains.
*/

/*
	 arr (sorted) = [9:00, 9:45, 9:55, 11:00, 15:00, 18:00]
		dep (sorted) = [9:20, 11:30, 11:50, 12:00, 19:00, 20:00]
*/
func platformsOptimal(arr []int, dep []int) int {
	var i, j = 0, 0
	var platforms = 0
	var result = 0
	sort.Ints(arr)
	sort.Ints(dep)
	for i < len(arr) && j < len(dep) {
		if arr[i] <= dep[j] {
			platforms++
			i++
		} else {
			platforms-- // you can check in there for negative value and reset to 0. But, result took max so anyway res is 0. if negative.
			j++
		}
		result = max(result, platforms)
	}
	return result
}
