package greedy

import (
	"fmt"
	"slices"
)

/*

You are given a set of N jobs where each job comes with a deadline and profit.
The profit can only be earned upon completing the job within its deadline.
Find the number of jobs done and the maximum profit that can be obtained.
Each job takes a single unit of time and only one job can be performed at a time.

Examples
Example 1:
Input: N = 4, Jobs = {(1,4,20),(2,1,10),(3,1,40),(4,1,30)}
Output: 2 60
Explanation: The 3rd job with a deadline 1 is performed during the first unit of time .
             The 1st job is performed during the second unit of time as its deadline is 4.
Profit = 40 + 20 = 60

Example 2:
Input: N = 5, Jobs = {(1,2,100),(2,1,19),(3,2,27),(4,1,25),(5,1,15)}
Output: 2 127

Explanation: The  first and third job both having a deadline 2 give the highest profit.
Profit = 100 + 27 = 127

*/

/*
 We need take jobs with maximium profit.  number of jobs to obtain this.
 - the slots = max(deadlines).
 in 1 th example we have 4 slots since max deadline is 4.
 in 2nt its 2.
  we need to take max possible profit in this slots.
  for example if deadline is 4. then jis job can be done in any 1,2,3,4 slot.
  and we can take it at the end.
  doing more urgent tasks.
 - pick most urgent and most profit job. at each step. to go greedy.
 1 take 40 with urdency 1. - since all other jobs deadline also 1.
 2 take 1 job, because only its lives not dedline exceeded.

*/
/*
 1. Brute force:
    - go i = 1 to N.
    - find all deadlines where deadline = i.
    - find max from them add to result, increment job count.
    return result.
    tc: N * N. Each time we need to check all jobs, for deadline and compare with max.
    sc O(1).
 2. How can we optimze deadline and max finding?
    - sort the array by deadline.
    - iterate over array go until a[i].deadline != a[i+1].deadline, and compare with max on each step.
    - if deadlines are not match, add max profit to result, and max = a[i+1].profit., increase counter of jobs.
    tc: nlogn+n - we sort and iterate once from given array.
    sc: O(1)

3. Usgin additional result array.
  - sort the array usign porfit.
  - find max deadline, x.
  - create additional array to store jobs, sceduled, with size x. and fill it with -1, means empty day.
  - iterate over sorted array.
	- job.deadline is last day.
	- if job is can done in th last day, mark the job for this day in array. and add profit to answer.
	- if it is not possible go prev days to find empty slot or day. todo this job.
	do it for each job.
 Why we need to sort the jobs by profit?
	- we need to maximize profit, this is mean we need first do job with max profit in deadline day.
		for example, if we have 2 jobs with 4 deadline and 20, 30 profit.
		if we fill, or do job with 20 profit first we loose 30 profit job in that day, or
			we need to od it one of prev days and it may affect to other prev day profits too, so its incorrect.
summary:
	Do first max profit jobs.
	Do it as possible in deadline day or close to deadline day.
with this 2 way we can got max profit from jobs.
*/

type job struct {
	Id       int
	Deadline int
	Profit   int
}

func maxprofit(jobs []job) int {
	//1. find max deadline
	x := jobs[0].Deadline
	for _, v := range jobs {
		x = max(x, v.Deadline)
	}
	//2. sort by profit in decreasing order.
	slices.SortStableFunc(jobs, func(a, b job) int {
		if a.Profit > b.Profit {
			return 0
		}
		return 1
	})

	fmt.Println(jobs)

	// fill schedule array.
	var schedule = make([]int, x)
	var res = 0
	for i, job := range jobs {
		for j := job.Deadline; j >= 0; j-- {
			if schedule[j] == 0 {
				schedule[job.Deadline] = i
				res += job.Profit
				break
			}
		}
	}
	return res
}
