package main

func main() {

}

func mySqrt(x int) int {
	low, high := 0, x
	for low <= high {
		mid := low + ((high - low) / 2)
		if mid*mid == x {
			return mid
		}

		if mid*mid > x {
			high = mid - 1
		} else if mid*mid <= x {
			low = mid + 1
		}
	}
	return high
}

/*
  1. Brute force.
   for i:=1 to i*i<=n; i++:
    if i*i <= n {
        ans = i
    }
    return ans
    but we got time limit in there.
    So what can we do?
 In there coming to mind is
 binary search on answers.
  so let's take 10
  i = 1 2 3 4 5 6 7 8 9 10
  low = 0
  high = x
  mid = low+high/2
    if mid*mid > x:
        then go left
    else if x< .. :
        then go to righ
    else:
        reutrn mid.
*/
