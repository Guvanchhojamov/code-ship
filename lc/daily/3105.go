package daily

// 3105. Longest Strictly Increasing or Strictly Decreasing Subarray
func main() {

}

func longestMonotonicSubarray(nums []int) int {
	var count = 1
	var maximum = 1 // minimum length of stricy subarray is 1.
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			count++
		} else {
			maximum = max(maximum, count) // we need to find longest.
			count = 1
		}
	}
	maximum = max(maximum, count)
	// iterate for second, for increasing order.
	// in there we cant set max=1, because or maximum myght be our answer.
	count = 1 // but we reset count to 1.
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			count++
		} else {
			maximum = max(maximum, count)
			count = 1
		}
	}
	maximum = max(maximum, count)
	return maximum // if no any stricy subarrays, we return 1.
}

// Test case 1:
/*
   [3,2,1]
   1. maximum = 3, count =3.
   2. maximum =3 count = 1.
   [1,4,3,3,2]
   1. [4,3], [3,2]
   2.  [1,4]
   result = 2.
*/

/*
 stricly increasing order.
 increasing order:
  a[i] itself -> increasing order.
  a[i] itselft -> decresing order too.

  increasing order:
 -  a[i] < a[i+1] <a[i+2]< ... < a[i+n] -> continuosly array like this is subaarray
   of stricly increasing order.
 -  a[i]>a[i+1]>a[i+2] > ... > a[i+3] - is continuosly subarray like this is
   subarray of stricly decreasing order.
  Key points:
   1. a[i] each element of array itself is:
      - increasing subarray with len=1
      - decreasing subarray with length=1
   2. We need to find longest subarray with that length.
   Brute force.
    iterate array 2 times and count for:
      1. increasing order.
      2. decreasing order
      3. if count ==0 then return 1 beacause any 1 elemmnt is
         increasing or in decreasing order.
    tc: O(n)+O(n) = O(n)
    sc: O(1)   is it ok?
   Can we optimize it?
   Maybe, but lets implement brute force first. Beacuse we have
   small constraint range in this problem.
   How can we use better algorythm in there?

*/
