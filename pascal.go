package main

import "fmt"

func pascal(row int, col int) int {
	if col == 0 || col == row {
		return 1
	}
	start := 1
	for i := 1; i <= col; i++ {
		start = start * (row + 1 - i) / i
	}
	return start
}

func main() {
	for i := 0; i < 50; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(pascal(i, j))
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
