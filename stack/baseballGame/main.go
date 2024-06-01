package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}

func calPoints(operations []string) int {
	var stack []int
	var result int
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "+":
			if len(stack) > 1 {
				result += stack[len(stack)-1] + stack[len(stack)-2]
				stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
			}
		case "D":
			result += stack[len(stack)-1] * 2
			stack = append(stack, stack[len(stack)-1]*2)
		case "C":
			if len(stack) > 0 {
				result -= stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
			}
		default:
			val, _ := strconv.Atoi(operations[i])
			stack = append(stack, val)
			result += val
		}
	}
	return result
}
