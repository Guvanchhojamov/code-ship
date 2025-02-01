package main

import "fmt"

/*
 Pattern-5: Inverted Right Pyramid
*/
func main() {
	var n = 6
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	p6Nums()
	p7()

}

func p6Nums() {
	var n = 6
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

/*
 00 01 02
 10 11 12
 20 21 22 23 24
*/
func p7() {
	var n = 6
	for i := 0; i < n; i++ {
		for j := 1; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
