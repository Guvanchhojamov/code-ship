package main

import (
	"fmt"
)

func printTable(n int) {
	// заголовка
	fmt.Print("\t")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println() //new line

	// разделител
	fmt.Print("\t")
	for i := 1; i <= n; i++ {
		fmt.Print("----\t")
	}
	fmt.Println()

	// умножения
	for i := 1; i <= n; i++ {
		fmt.Printf("%d |\t", i) // левая колонка
		for j := 1; j <= n; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println() // to new line
	}
}

func main() {
	printTable(5)
}
