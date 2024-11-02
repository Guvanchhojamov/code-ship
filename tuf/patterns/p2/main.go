package main

import "fmt"

/*
  Pattern-2: Right-Angled Triangle Pattern
*/

/*
 00
 10 11
 20 21 22
*/
func main() {
	var n = 6
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j {
				fmt.Printf("%v%v ", i, j)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	p2Nums()
	p3Nums()
}

func p2Nums() {
	var n = 6
	var count = 0
	for i := 0; i < n; i++ {
		count = 0
		for j := 0; j < n; j++ {
			if i >= j {
				count += 1
				fmt.Printf("%d", count)
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}

func p3Nums() {
	var n = 6
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if i >= j {
				fmt.Printf("%d", i)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
