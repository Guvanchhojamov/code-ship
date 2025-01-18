package main

// 278. First Bad Version
func main() {

}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	var low, high = 1, n
	var res = n
	for low <= high {
		mid := low + ((high - low) / 2)
		if isBadVersion(mid) { // leetcode own function
			res = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}

/*
for i:=0;i<=n;i++ {
        if isBadVersion(i) {
            return i
        }
    }
    return n
*/

/*
  Need to solve with bs
  what main problem syas?
  Find bad first bad version from given versions
  [1..n]

  All bad versions after first bad, is bad too
    For example
    [1,2,4,5,6,7,8,9,10]
    if 5 is firstBad version
       5,6,7,8,9,10 all versions is bad.
       we need to find first one.

    So with binary seach we can find.
       Need to find smallest index where `IsBad`
         Lower bound:
         the smallest number bigger or equal to target.
so,
     for i-> 1 to n:
   the smallest i where  i >= target -->   or target is true.
    the smalles i where  i == true.
    let's implement lower bound.
*/
