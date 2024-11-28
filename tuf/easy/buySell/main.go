package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	res := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(res)
}

/*
Problem Statement:
You are given an array of prices where prices[i] is the price of a given stock on an ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
*/

/*
  1. keep min price  default 0
  2. keep maxProfit  default a[0]

   3. if a[i] < minPrice then minPrice=a[i]
   4. if a[i]-min> maxProfit  then maxProfit = a[i]-min

 return maxProfit.

*/
// Optimal aproach
func maxProfit(prices []int) int {
	var maxProfit = 0
	var minPrice = prices[0]

	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxProfit = max(maxProfit, price-minPrice)
	}

	return maxProfit
}

/*
   [7,1,5,3,6,4]    res = 5  [1-buy, 6-sell]

   maxProfit = 5
   minPrice = 1


*/