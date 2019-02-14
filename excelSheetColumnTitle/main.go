package main

import (
	"fmt"
)

func main() {
	fmt.Println(convertToTitle(27))
}

func convertToTitle(n int) string {
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	result := ""
	for n > 0 {
		n -= 1
		mole := n % 26
		n = (n - mole) / 26
		result = letters[mole] + result
	}
	return result
}
