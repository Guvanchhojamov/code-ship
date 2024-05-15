package main

import "fmt"

func main() {
	fmt.Println(minOperations([]string{"d1/", "d2/", "../", "d21/", "./"}))
}

func minOperations(logs []string) int {
	var level int
	for _, log := range logs {
		switch log {
		case "./":
			continue
		case "../":
			if level > 0 {
				level -= 1
			}
		default:
			level += 1
		}
	}
	return level
}
