package stack

/*
901. Online Stock Span~
Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.

The span of the stock's price in one day is the maximum number of consecutive days (starting from that day and going backward) for which the stock price was less than or equal to the price of that day.

For example, if the prices of the stock in the last four days is [7,2,1,2] and the price of the stock today is 2, then the span of today is 4 because starting from today, the price of the stock was less than or equal 2 for 4 consecutive days.
Also, if the prices of the stock in the last four days is [7,34,1,2] and the price of the stock today is 8, then the span of today is 3 because starting from today, the price of the stock was less than or equal 8 for 3 consecutive days.
Implement the StockSpanner class:

StockSpanner() Initializes the object of the class.
int next(int price) Returns the span of the stock's price given that today's price is price.


Example 1:

Input
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
Output
[null, 1, 1, 1, 2, 1, 4, 6]

Explanation
StockSpanner stockSpanner = new StockSpanner();
stockSpanner.next(100); // return 1
stockSpanner.next(80);  // return 1
stockSpanner.next(60);  // return 1
stockSpanner.next(70);  // return 2
stockSpanner.next(60);  // return 1
stockSpanner.next(75);  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
stockSpanner.next(85);  // return 6
*/

type StockSpanner struct {
	Idx      int
	Prices   []int
	PrevMaxi []int
}

func Constructor2() StockSpanner {
	return StockSpanner{
		Idx:      0,
		Prices:   []int{},
		PrevMaxi: []int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	var count = 0
	for len(this.PrevMaxi) > 0 &&
		len(this.Prices) > 0 &&
		this.Prices[this.PrevMaxi[len(this.PrevMaxi)-1]] <= price {
		this.PrevMaxi = this.PrevMaxi[:len(this.PrevMaxi)-1]
	}
	if len(this.PrevMaxi) == 0 {
		count = this.Idx + 1
	} else if len(this.Prices) > 0 {
		count = this.Idx - this.PrevMaxi[len(this.PrevMaxi)-1]
	}
	this.Prices = append(this.Prices, price)
	this.PrevMaxi = append(this.PrevMaxi, this.Idx)
	//  fmt.Println(this.Prices, this.PrevMaxi, this.Idx-1)
	this.Idx++

	return count
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

/*
  We need return for given day price, <= continues prices, before current day.
brute force:
- implement stock with: prices = []int, count = int.
- on each next call, add price to prices. go until
prices[i]>price and count else return count. and reset count.
tc: ~n*n
sc: n for list.
2. Optimize with monothonic stack.
- prices array -> containts prices.
- prevMaxi array -> stack contains prev greater element idx.
for price>=prices[last]: st[:-1].
count = priceIdx - st[last].
return count.
*/
