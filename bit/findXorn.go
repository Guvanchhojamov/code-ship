package bit

/*
find xor of nums 1 to n
*/
func xorToN(n int) int {
	// bf is 0 to n make xor. tc:n
	/*
	   1^1 = 1
	   1^2 = 3
	   1^2^3 =0
	   1^2^3^4 = 4
	*/
	if n%4 == 1 {
		return 1
	}
	if n%4 == 0 {
		return n
	}
	if n%4 == 2 {
		return n + 1
	}
	if n%4 == 3 {
		return 0
	}
	return 0
}

/*
firn xor of range L and R
*/
func xorRange(l, r int) int {
	lx := xorToN(l - 1)
	rx := xorToN(r)
	res := lx ^ rx
	return res
}

/*
for example if l=4,r=7
 (1^2^3) ^ (1^2^3^4^5^6^7) = 0^4^5^6^7

*/
