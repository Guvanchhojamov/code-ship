package main

import "fmt"

func main() {
	res := xorAllNums([]int{2, 1, 3}, []int{10, 2, 5, 0})
	fmt.Println(res)
}

func xorAllNums(nums1 []int, nums2 []int) int {
	var xn1, xn2, res int
	for _, num := range nums1 {
		xn1 ^= num
	}
	for _, num2 := range nums2 {
		xn2 ^= num2
	}
	for _, n1 := range nums1 {
		if len(nums2)%2 == 0 {
			res ^= 0 ^ xn2
		} else {
			res ^= n1 ^ xn2
		}
	}
	return res
}

// Optimal solution:

/*
  2^10^2^2^2^5^2^0

shuny sheyle uytgedip bolya:
   2^2^2^2^10^2^5^0 kopeltmek tabliza yaly.
   2^2^2^2 -> hemishe 0 eger jubut bolsa
   2^2^2^2^2 -> hemishe sho sanyn ozi eger tak bolsa
   2^2^2^2^10^2^5^0 == 0^10^2^5^0 -> ikisi den
   diymek
    0^XOR(nums2-lerin Xor-y)  eger jubut bolsa
    2^XOR(nums2-lerin Xor-y)   eger tak bolsa

  xn1 int
  xn2 int
   for i:=to nums1:
     xn1^=nums1[i]

   for j:=to nums2:
     xn2^=nums2[i]

  for i:= to nums1:
     if len(nums2) % 2:
       res^=0^xn2
     else if len(nums2) %2 != 0:
       res^=nums[i]^xn2

  return res

  tc: 2n1+n2 -> O(2n) -> O(n)
  sc: O(1)
  lets try  this one
*/

/*
 brute force:
 num3[]
  for i->nums1:
   for j -> nums2:

     nums3[i] = append.n[1]^n[2]

 for i-> nums3:
    res^=nums3[i]

 return res

 TC: n*n+ n*n => 2n*n -> O(n*n)
 SC: n*n -> O(n*n)

   shu yerde nums3-i ulanman chozup bileris.
    for i->nums1:
        for j -> nums2:

     res ^= nums1[i]^nums2[j]

    return res

 TC: n*n => O(n*n) bolar
 SC: O(1) additional memory ulanylanok.
*/
