package main

import "fmt"

/*
844. Backspace String Compare

Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.
Note that after backspacing an empty text, the text will continue empty.

*/

func main() {
	fmt.Println(backspaceCompare("y#fo##f", "y#f#o##f"))
}

func backspaceCompare(s string, t string) bool {
	return clean(s) == clean(t)
}

func clean(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '#':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
