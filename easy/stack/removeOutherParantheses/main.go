package main

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
}
func removeOuterParentheses(s string) string {
	var stack, result string
	var opens, closes int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack += string(s[i])
			opens++
		case ')':
			stack += string(s[i])
			closes++
		}
		if opens == closes {
			if len(stack) > 1 {
				stack = stack[1 : len(stack)-1]
			} else {
				stack = ""
			}
			opens = 0
			closes = 0
		}
		result += stack
		stack = ""
	}
	return result
}
