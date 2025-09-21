package greedy

import (
	"sort"
)

/*

435. Non-overlapping Intervals
Given an array of intervals intervals where
    intervals[i] = [starti, endi]
return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Note that intervals which only touch at a point are non-overlapping.
    For example, [1, 2] and [2, 3] are non-overlapping.


Example 1:
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.

Example 2:
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.

Example 3:
Input: intervals = [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.


Constraints:

1 <= intervals.length <= 10^5
intervals[i].length == 2
-5 * 10^4 <= starti < endi <= 5 * 10^4

*/
/*
 - given intervals, we need to find count of intervals removed, to do rest of other intervals are non overlaping.
 - if no overlaps result is 0.
 - one touch is non overlap. 1,2; 2,3 - 2:2 not overlaping.

intervals = [[1,2],[2,3],[3,4],[1,3]].
1,2  2,3  3,4  1,3

in there 1,2 and 1,3 - are overlapping.
    if we remove 1,3 all other intervals are not overlapping. res: 1
[[1,2],[1,2],[1,2]]
1,2; 1,2; 1,2
 to do correct we need to remove any of 2 overlaps. res =2
 [[1,2],[2,3]]; res = 0.
*/
/*
  - how we can find overlap array. and we must remove minimum number of overlaps.
  to got overlaps nearest, overlaps we need to move them sequantally.
  to move near lets sort the intervals, by end time.
  1,2  2,3  3,4  1,3
  after sort
  1,2  2,3  1,3  3,4
  in there we can see exacty what is overlaping?
    now for each interval check:
        if i.end > j.start:
        count++
    why this work?
     because we sort by end date.
    so, after sort, i.end date is < i+1.end date always, and if i.end > i+1.start this means i+1.starts before i.ends. and this is overlap.
tc: nlogn + n
sc: 1
[[1,100],[11,22],[1,11],[2,12]]
  1,11  2,12  11,22  1,100

*/

func eraseOverlapIntervals(intervals [][]int) int {
	var nonOverlaps = 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	//fmt.Println(intervals)
	var lastEnd = intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] >= lastEnd {
			nonOverlaps++
		}
	}
	return len(intervals) - nonOverlaps
}
