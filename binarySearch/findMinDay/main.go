package main

import "slices"

func main() {

}

/*
You are given an integer array bloomDay, an integer m and an integer k.
You want to make m bouquets. To make a bouquet, you need to use k adjacent flowers from the garden.
The garden consists of n flowers, the ith flower will bloom in the bloomDay[i] and then can be used in exactly one bouquet.
Return the minimum number of days you need to wait to be able to make m bouquets from the garden. If it is impossible to make m bouquets return -1.
*/

/*
Example 1:

Input: bloomDay = [1,10,3,10,2], m = 3, k = 1
Output: 3
Explanation: Let us see what happened in the first three days. x means flower bloomed and _ means flower did not bloom in the garden.
We need 3 bouquets each should contain 1 flower.
After day 1: [x, _, _, _, _]   // we can only make one bouquet.
After day 2: [x, _, _, _, x]   // we can only make two bouquets.
After day 3: [x, _, x, _, x]   // we can make 3 bouquets. The answer is 3.

Example 2:
Input: bloomDay = [1,10,3,10,2], m = 3, k = 2
Output: -1
Explanation: We need 3 bouquets each has 2 flowers, that means we need 6 flowers. We only have 5 flowers so it is impossible to get the needed bouquets and we return -1.
Example 3:

Input: bloomDay = [7,7,7,7,12,7,7], m = 2, k = 3
Output: 12
Explanation: We need 2 bouquets each should have 3 flowers.
Here is the garden after the 7 and 12 days:
After day 7: [x, x, x, x, _, x, x]
We can make one bouquet of the first three flowers that bloomed. We cannot make another bouquet from the last three flowers that bloomed because they are not adjacent.
After day 12: [x, x, x, x, x, x, x]
It is obvious that we can make two bouquets in different ways.


Constraints:

bloomDay.length == n
1 <= n <= 105
1 <= bloomDay[i] <= 109
1 <= m <= 106
1 <= k <= n

*/

/*
  [7 7 7 7 11 13 7 7 ] m=2 k=3
  We need m buqtets with k item in.
  Items can be day,by bay near of each other,
  Example, for make 1 buqet
   days must be near: 1 2 3 -> in all days we need can make buqets.
   1. Brute force.
   key points:
    1 2 3 4 5 6 7 8 9
    min   to maxArr -> this is maxDay.
    if mxk > len(a) return false, because we can not make more than n buqets.
    We need helper function like:
     for each day check is it possible or not
     if for all  days is not possible return -1.

    for i:= minArr-> maxArr:
        if isPossible(arr, i, m , k) {
            return i
        }
    return -1

    func isPossible(arr, day, m, k ){
        possItems:=0
        buqesCount = 0
        for i:=range arr:
            if day % arr[i] == 0 {
                possibleItems++
            }else {
                buqesCount+=possibleItems / m
                possibleItems = 0
            }
            buqesCount += possibleItems / m
        return m == buqesCount
    }
*/

func minDays(bloomDay []int, m int, k int) int {
	var minDay = slices.Min(bloomDay)
	var maxDay = slices.Max(bloomDay)

	for day := minDay; day <= maxDay; day++ {
		if isPossible(bloomDay, day, m, k) {
			return day
			break // we need minumim day here, so we can break.
		}
	}
	return -1
}

func isPossible(blooms []int, day int, m, k int) bool {
	var possCount = 0
	var buqesCount = 0
	for _, bloomDay := range blooms {
		if day >= bloomDay {
			possCount++
		} else if possCount > 0 {
			buqesCount += possCount / k
			possCount = 0
		}
	}
	buqesCount += possCount / k
	return buqesCount == m
}

/*
  [7, 7, 7, 7, 13, 11, 12, 7 ] m=2 k=3     res = 13.
   day =12
   pC=4
   bC=1+1=2
   2=2 return true.
*/
