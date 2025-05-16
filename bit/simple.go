package bit

import "fmt"

/*
We try to solve simple problems, related with bit manipulations.
1. Swap two numbers
2. Check if ith  bit is set or not
3. set ith Bit
4. clear ith Bit
5. Toggle ith Bit
6. Remove the last set Bit
7. Check if a number is power of 2 or not
8. Count the number of set bits
*/

func main22() {
	fmt.Println("Hello, World!")
	swapNums(13, 5)
	isSetIthBit(13, 1) //1101
	setIthBit(13, 1)
	clearIthBit(13, 3)
	toggleIthBit(13, 1)
	removeLastSetBit(13)
	checkNumPow2(64)
	countNumberOfSetBits(13)
}

func countNumberOfSetBits(x int) {
	// 13 1101
	/*
			   bf: conv to bin, count 1's and return.
			   ok, we can 0->31  each bit with and if it is set count++
			   and return count.
		       tc:
	*/
	var cnt = 0
	for i := 0; i < x+1; /*or 32*/ i++ {
		if x&(1<<i) > 0 {
			cnt++
		}
	}
	fmt.Println("set bit count: ", cnt)
	// other solution is we can check each bit shifting to right until n==0
	cnt = 0
	for x > 0 {
		if x&1 == 1 {
			cnt++
		}
		x = x >> 1
	}
	fmt.Println("caluculated with shift: ", cnt)
}

func checkNumPow2(x int) {
	/*
			   bf: 0 to max_int check
		       all 2^x nums binary are 100..0
		       all 2^x-1 nums bin are  011..1
		       so if we and this vals we got 0
		       if num is pow of 2 the result of
		       n&n-1 must be 0.
	*/
	if x&x-1 == 0 {
		fmt.Println(x, "-is power of 2")
	} else {
		fmt.Println(x, "-is not power of 2")
	}
	// second approach si using not x-1= 0111 -> ^(x-1) = 1000 so x and NOT(x-1) ==0 then return true.
}

func removeLastSetBit(x int) {
	/*
			   12 = 1100
		       we need to remove last bit
		       bf: iterate from back iuntil i==1, set=0 break.
		       how we can check set or not ith bit? We already does it, so
		       for int it is from 0 ot 31.
		        check each bit is set:
		            if set make it 0 and break.
	*/
	fmt.Printf("before: %04b\n", x)
	var res = x
	for i := 0; i <= 31; i++ {
		if x&(1<<i) > 0 { // then it is the first set bit,so then toggle it
			res = x &^ (1 << i)
			break
		}
	}
	fmt.Printf("after remove last set: %04b\n", res)
	fmt.Printf("simple remove last set: %04b\n", x&x-1)
	// The simplest way is x-1 is always not set in last indx :).
	// x&x-1 is last set removed element.
}

func toggleIthBit(x, i int) {
	/*
	   toogle means set if not set, not set if set
	   1101 i=2
	   0100 = 1<<2
	   xor we can toggle the ith bit
	   1001
	   i=1 1101^0010 = 1111
	*/
	fmt.Printf("before: %04b\n", x)
	res := x ^ (1 << i)
	fmt.Printf("after toggle: %04b\n", res)
}

func clearIthBit(x, i int) {
	// clear means not set.
	/*
			   1101 i=2
			   1<<2 = 0100; num and+ not(1<<i) it clears the ith bit.
		       1<<2 = 0100 -> ^(0100)= 1011
		       1101 & 1011 => 1001 -> cleared 2nd bit
	*/
	fmt.Printf("before: %04b\n", x)
	res := x & ^(1 << i)
	fmt.Printf("after clear: %04b\n", res)
}

func setIthBit(x, i int) {
	// 1101 i=1
	// 1<<1 = 0010
	// num or 1<<i = sets ith bit
	fmt.Printf("before: %04b\n", x)
	res := x | (1 << i)
	fmt.Printf("after set: %04b\n", res)
}

func isSetIthBit(x int, i int) {
	//bf: conv to bin [] check i-th bit.
	/*
	   1101  i=2
	   if 1<<i and num >0 then i-th bit is set. else non set
	*/
	if x&(1<<i) > 0 {
		fmt.Println("set")
	} else {
		fmt.Println("not set")
	}
}

func swapNums(a, b int) {
	fmt.Println("before: ", a, b)
	a = a ^ b
	b = a ^ b // (a^b)^b
	a = a ^ b // a^(a^b)
	fmt.Println("after swap: ", a, b)
}
