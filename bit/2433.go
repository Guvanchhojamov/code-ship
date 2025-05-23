package bit

/*
2433. Find The Original Array of Prefix Xor
You are given an integer array pref of size n.
Find and return the array arr of size n that satisfies:

pref[i] = arr[0] ^ arr[1] ^ ... ^ arr[i].
Note that ^ denotes the bitwise-xor operation.

It can be proven that the answer is unique.

Example 1:

Input: pref = [5,2,0,3,1]
Output: [5,7,2,3,2]
Explanation: From the array [5,7,2,3,2] we have the following:
- pref[0] = 5.
- pref[1] = 5 ^ 7 = 2.
- pref[2] = 5 ^ 7 ^ 2 = 0.
- pref[3] = 5 ^ 7 ^ 2 ^ 3 = 3.
- pref[4] = 5 ^ 7 ^ 2 ^ 3 ^ 2 = 1.
Example 2:

Input: pref = [13]
Output: [13]
Explanation: We have pref[0] = arr[0] = 13.
Constraints:
1 <= pref.length <= 10^5
0 <= pref[i] <= 10^6
*/
/*
Input: pref = [5,2,0,3,1]
Output: [5,7,2,3,2]
Explanation: From the array [5,7,2,3,2] we have the following:

We are given the pref array, so we need return the original array from this.
- pref[i] = orig[0]^orig[1]^orig[2]... orig[i]
- constraints are big.
- result output is unique.
pref = [5,2,0,3,1]
out=[x,y,z]
0^5 = x
0^5^2 = y
0^5^2^0 = z
so like this we can find the elements of output.
We use prefixXor method,
 bf is:
-create output empty arr with len = pref.
-for each output[i]:
	iterate nums 0->i xor:=xor^pref[i]
	add this xor to output arr.
return output arr.
tc: O(N*N) for each i in pref N and it is increasing, so for last element is N. worth case O(n*n).
sc: O(1) we dont use any extra space.

2. We can optimize the xor calculation with prefix_xor. How we can see
pref = [5,2,0,3,1]
out=[x,y,z]
0^x = 5
0^x^y = 2
0^x^y^z = 0
out[0] = pref[0]
out_xor = 0^pref[i]
for i=1 to len(pref):
 out[i] = pref[i]^out_xor
 out_xor = out_xor ^ out[i]
return out.
tc: O(n) we iterate one time pref arr .
sc: O(1) not used any extra space.
pref = [5,2,0,3,1]
out =  [5,7,2,0,0]; out_xor = 0^5=5
i=1
out[i] = 2^5 = 7
out_xor = 5^7 = 2
i=2
out[i]=0^2
out_xor = 2^2 =0
*/

func findArrayOld(pref []int) []int {
	var out = make([]int, len(pref))
	out[0] = pref[0]
	var out_xor = 0 ^ out[0]
	/*
	   5,7,2
	   5,x,y

	*/
	for i := 1; i < len(pref); i++ {
		out[i] = out_xor ^ pref[i]
		out_xor = out_xor ^ out[i]
	}
	return out
}

/*
[5,2,0,3,1]
*/

func findArray(pref []int) []int {
	for i := len(pref) - 1; i > 0; i-- {
		pref[i] = pref[i] ^ pref[i-1]
	}
	return pref
}
