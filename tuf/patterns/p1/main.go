package main

import "fmt"

/*
  Given an integer N, print the following pattern.
  Example 1:
Input: N = 3
Output:
* * *
* * *
* * *
*/
func main() {
	p1(6)
}

func p1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%s ", "*")
		}
		fmt.Println()
	}
}
