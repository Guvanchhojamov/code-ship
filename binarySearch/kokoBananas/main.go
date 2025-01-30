package main

import "math"

func main() {

}

/*
Koko loves to eat bananas.
There are n piles of bananas, the ith pile has piles[i] bananas.
The guards have gone and will come back in `h` hours.

Koko can decide her bananas-per-hour eating speed of `k`.
Each hour, she chooses some pile of bananas and eats `k` bananas from that pile.
If the pile has less than `k` bananas,
she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer `k` such that she can eat all the bananas within `h` hours.

Example 1:
Input: piles = [3,6,7,11], h = 8
Output: 4
Example 2:

Input: piles = [30,11,23,4,20], h = 5
Output: 30
Example 3:

Input: piles = [30,11,23,4,20], h = 6
Output: 23

Constraints:
1 <= piles.length <= 104
piles.length <= h <= 109
1 <= piles[i] <= 109

*/

/*
  piles [1..n]
  banannas in each piles X
  h = hours witch koko can eat  bananas.
  k = banana per hours need to eat for koko

  Problem, Find minimum K such that, koko must eat all bananas in H hours.




  1. Brute force:

     The max hour is  max(piles), = n
     k = 1
     totalHour = 0
     for 1 -> piles:
        if func(piles, k) <= h:
            return h
     return max(piles)


    func(piles, k) int :
        for i:=range piles:
            totoalHour +=ceil(piles[i] / k)
    return totalHours
*/

func minEatingSpeed(piles []int, h int) int {
	var maxPile = maxInPile(piles)

	for i := 1; i <= maxPile; i++ {
		if totalHours(piles, i) <= h {
			return i
		} else {
			break
		}
	}
	return maxPile
}

func totalHours(piles []int, k int) int {
	var totalHour int
	for i := range piles {
		totalHour += int(math.Ceil(float64(piles[i]) / float64(k)))
	}
	return totalHour
}

func maxInPile(piles []int) int {
	var maximum = math.MinInt

	for _, val := range piles {
		maximum = max(maximum, val)
	}
	return maximum
}

/*
  This is the brute force approach, and we got time complexity int there
So we need to optimize it with using binary search approach.
*/

func minEatingSpeedBs(piles []int, h int) int {
	var maxPile = maxInPile(piles)
	var low, high = 1, maxPile
	// var ans = maxPile
	for low <= high {
		mid := (low + high) / 2
		wastedHours := totalHours(piles, mid)
		if wastedHours <= h {
			high = mid - 1
			//          ans = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
