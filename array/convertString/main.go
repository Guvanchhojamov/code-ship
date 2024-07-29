package main

import "fmt"

func main() {
	minimumMoves("OXOX")
}

// lc. 2027
/*
XXOX -> OOOX

1. find len 3, 2, 1

*/
func minimumMoves(s string) int {
	reqLen := 3
	res := 0
	c := []byte(s)
	for reqLen > 0 {
		xCount := 0
		oCount := 0
		k := 0
		for i := 0; i < len(c); i++ {
			k++
			if c[i] == 'X' {
				xCount++
			}
			if c[i] == 'O' {
				oCount++
			}

			if i > 1 {
				fmt.Println("xCount", xCount, "k", k)
				if k == 3 || i-k < len(c) {
					if xCount == reqLen {
						c[i] = 'O'
						c[i-1] = 'O'
						c[i-2] = 'O'
						res++
					}
					xCount = 0
					k = 0
				}
			}
		}

		fmt.Println(string(c), res, reqLen)
		reqLen--
	}
	return 0
}
