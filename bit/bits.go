package bit

import "fmt"

/*
   In there some simple bit manipulation problems.
What is the bit manipulation?
It is th change bits from num
Comp uses only 2 binary.
x32 or x64
0 - sign bit is unsigned number.
1 - sign bit is signed number.
if sign bit==1 then num<0
if sign bit == 0 then num > 0.
*/

/*
1. Binary number conversion + (1st, 2nd compliment)
	- the 1st compliment is convert num to 2 binary representation.
    - the 2nd compliment is add 1 to 2 binary conversion.
2. Operators (and,or,xor,not,shift)
3. Swap 2 number  (using bits.)
4. Check if the ith bit set or not.
5. Extract the ith bit.
6. Set the ith bit.
7. clear the ith bit.
8. toggle the ith bit.
*/

/*
 1. binary conversion is mod to 2 until n == 0.
    return reversed results.
    while n>0:
    c:=n%2
    res = c+res  // reversed
    n = n/2
    revers(res) return.

 2. convert from binary to 10 th
    what is the formula?
    0011 is
    start from back and ch*2^i

and sum of all nums then.

		  0 0 1 1
		s= 1*2^0+1*2^1+0*2^2+0*2^3
		s = 1+2 = 3
	 how we can do it in program?
*/
func toBin(n int) []int {
	d := fmt.Sprintf("%04b", n) // simple
	fmt.Println(d)
	var res []int
	for n > 0 {
		c := n % 2
		res = append(res, c)
		n = n / 2
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func toDecimal(s string) int {
	c := []byte(s)
	res := 0
	for i := len(c) - 1; i >= 0; i-- {
		p := getPowof2(i)
		num := int(c[i] - '0')
		res = res + num*p
	}
	return res
}

func getPowof2(x int) int {
	var r = 1
	if x == 0 {
		return 1
	}
	for i := 1; i <= x; i++ {
		r = r * 2
	}
	return r
}

/*
this is the brute force solutions.
But we have some formula for example go get the decimal of binary number.
 1. How we can get, power in easy way?
*/
func toDecimal2(s string) int {
	c := []byte(s)
	res := 0
	var p2 = 1
	for i := len(c) - 1; i >= 0; i-- {
		num := int(c[i] - '0')
		if num == 1 {
			res = res + p2
		}
		p2 = p2 * 2
	}
	return res
}

/* we can update like this taking only 1-th, and calculate 2 pow like prefix multiple */

/*
  1 th compliment is convert num to bin, and flip all binaries.
	on flip 1 -> 0 and 0->1
  2 th compliment is add +1 to flipped array.
flip:
	iterate over s:
	if 0 make it 1.
	if 1 make it 0.
add +1:
	iterate over s starting from back:
	if 0 make it 1 and break.
	if 1 make it 0 and continue.
	in the end check if s[0] == 0 then add this 1 like as external.
	1+s because 1 is remaining after all iterations.
*/
