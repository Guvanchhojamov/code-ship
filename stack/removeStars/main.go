package main

func main() {
	removeStars("leet**cod*e")
}

func removeStars(s string) string {
	var stack []rune
	for _, v := range s {
		if v == '*' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, v)
	}
	return string(stack)
}
