package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) > 1 {
			stack = checkAndRemove(stack)
		}
	}
	return string(stack)
}
func checkAndRemove(stack []byte) []byte {
	if stack[len(stack)-2] == stack[len(stack)-1] {
		stack = stack[0 : len(stack)-2]
	}
	return stack
}
