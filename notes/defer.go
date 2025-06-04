package notes

/*

https://victoriametrics.com/blog/defer-in-go/
 defer has  3 type.
 - heap-allocated defer.
 - stack-allocated defer.
 - open-code defer.
The difference is where the defer object is allocated, as name says.
before go-1.13 ve only has heap-allocated defer.
After 1.14 added 2 more stack and open-code.
func testDefer(a int) {
	if a == unpredictableNumber {
		defer println("Defer in if") // stack-allocated defer
	}
	if a == unpredictableNumber+1 {
		defer println("Defer in if") // stack-allocated defer
	}

  for range a {
    defer println("Defer in for") // heap-allocated defer
  }
}
 - if defer in for loop it is heap-allocated, because the size is not fixed.
 - otherwise defer is stack-allocated.
 - to be open-coded stack the number of defers must be  less< 8.
		because the open-coded method under the hood managed with bitmask.
*/
