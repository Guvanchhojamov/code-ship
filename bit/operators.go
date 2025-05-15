package bit

import "fmt"

/*
For bit manipulation we have some operators,
and,or,xor,not,shift left, shift right.
and = &&  both or not.
or  = || 1 of both
xor = must 1 of both.
not = ! in reverse manner.
shift we have 2 shifts left and right.
shift means 1 bit left or right.
<< - left shift 1 bit
>> - right shift 1 bit.
*/

/*
  move forwar to bit manipulation playlist.

*/

func bitAnd(x, y int) {
	res := x & y
	fmt.Println("bit and is: ", res)
	/*
	   1101
	   0111
	   0101 = 5
	   so returns 5
	*/
}

func bitOr(x, y int) {
	res := x | y
	fmt.Println("bit or is: ", res)
	/*
	   13 | 7
	   1101 | 0111 = 1111  = 15
	*/
}

func bitXor(x, y int) {
	res := x ^ y
	fmt.Println("bit xor is: ", res)
	/*
	   xor is just 1 need to true, exatly one not both.
	   1101 ^ 0111 = 1010 = 10
	*/
}

func bitNot(x int) {
	res := ^x
	fmt.Println(res)
	/*
	    how to get bitnot?
	   1. flip the bin of number
	   2. check sing
	       signed: -> return 2'th compliment
	       unsigned: -> stop, and return flipped dec.
	*/
}
func shiftLeft(x int) {
	res := x << 1 // 1 bit to left
	fmt.Println("left shifted is: ", res)
	/*
	   13
	    1101 << 1 = 1101+0 = 26
	*/
}

func shiftRight(x int) {
	res := x >> 1
	fmt.Println("one bit to right is:", res)
	/*
	   1101 >>1 = '1'-'110' = 6
	*/
}

func shiftKright(x, k int) {
	// we have a formaula.
	// n>>k = n/2^k
	// 2^k we use math.Pow(2,x) or
	// simple shift bit to left
	res := x / powof2(k)
	fmt.Println("shift k right is:", res)
}

func powof2(k int) int {
	return 1 << k
}
