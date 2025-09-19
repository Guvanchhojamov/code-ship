package greedy

import (
	"fmt"
	"sort"
)

/*
You are given timings of n meetings in the
	form of (start[i], end[i]) where start[i] is the start time of meeting i and end[i]
	is the finish time of meeting i.
Return the maximum number of meetings that can be accommodated in a single meeting room,
when only one meeting can be held in the meeting room at a particular time.

Note: The start time of one chosen meeting can't be equal to the end time of the other chosen meeting.

Examples :
Input: start[] = [1, 3, 0, 5, 8, 5], end[] =  [2, 4, 6, 7, 9, 9]
Output: 4
Explanation: Maximum four meetings can be held with given start and end timings. The meetings are - (1, 2), (3, 4), (5,7) and (8,9)

Input: start[] = [10, 12, 20], end[] = [20, 25, 30]
Output: 1
Explanation: Only one meetings can be held with given start and end timings.

Input: start[] = [1, 2], end[] = [100, 99]
Output: 1

Constraints:
1 ≤ n ≤ 105
0 ≤ start[i] < end[i] ≤ 106

*/

/*
  given meeting start and end times, and we have 1 room.
  we need to find, how many meeting we can accommodated.
  - end[i] and start[i+1] != always.

  Input: start[] = [1, 3, 0, 5, 8, 5], end[] =  [2, 4, 6, 7, 9, 9]
  1,2
  3,4
  0,6 x
  5,7
  8,9
  5,9 x
- in this case we van 4 meetings.
 so, we need to go with greedy approach.
	since,
	if end[i]< end[j] -> then start[i] < start[j] always.
We need ti find shortest meetings first then go to the next.
	start from most shortest, and go to most longest.
 in this scenario we got max number of meeting.
because in otherways like example 0,6 and 1,2,3,4.
one meetign 0,6 closes all others and we lost 3 meetings.
so, How we can go from short to long?
- if we sort using start time,
then metting can start early and can finnish at the end of the day, so its worng.
- if we sort using end time.
	then if finishes earlier then, its short, and we have more time to other meetings.
	if we can precess more meetigs we can got max number of meetings.
approach:
	1. create Meeting: startTime, endTime.
	2. create meetings: []Meeting
	iterate over start or end and fill meetings.
	3. sort meetings array using sortFund and sort by endDate.
	4. iterate over sorted meetings array and:
		start with count=1, we can process at least 1 meeting.
		for i: 1->n
		if meeting[i-1].end < meeting[i].start:
			count++
	5. at the end return count.
tc: n - for fill + nlong - for sort. + n for sorted. n+nlogn.
sn: n for meetings array.
*/

func nMeetings(start, end []int) int {
	var count = 1
	type Meeting struct {
		StartTime int
		EndTime   int
	}
	var meetings = []Meeting{}

	for i := 0; i < len(start); i++ {
		meetings = append(meetings, Meeting{StartTime: start[i], EndTime: end[i]})
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].EndTime < meetings[j].EndTime
	})

	fmt.Println(meetings)
	endTime := meetings[0].EndTime
	for i := 1; i < len(meetings); i++ {
		if endTime < meetings[i].StartTime {
			count++
			endTime = meetings[i].EndTime
		}
	}

	return count
}
