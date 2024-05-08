package main

import "fmt"

func main() {
	fmt.Println(numberDeclention(5))
}

func numberDeclention(n int) string {
	lastDigit := n % 10
	lastTwoDigits := n % 100

	if lastTwoDigits >= 11 && lastTwoDigits < 14 {
		return fmt.Sprintf("%d компютеров", n)
	}
	switch lastDigit {
	case 1:
		return fmt.Sprintf("%d компютеров", n)
	case 2, 3, 4:
		return fmt.Sprintf("%d компютера", n)
	default:
		return fmt.Sprintf("%d компютеров", n)

	}
}
