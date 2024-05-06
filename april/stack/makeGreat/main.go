package main

import (
	"fmt"
)

const requiredDiff = 32

func main() {
	fmt.Println(makeItGreat("abBAcC"))
}

func makeItGreat(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) > 1 {
			if stack[len(stack)-2]-stack[len(stack)-1] == requiredDiff {
				stack = stack[0 : len(stack)-2]
			}
		}
	}
	return string(stack)
}
