package main

import (
	"fmt"
	"strconv"
)

const array_len = 8

func main() {
	blanks := ""
	for i := 0; i < array_len*2; i++ {
		blanks += " "
	}
	arr := generate(array_len)
	for index, value := range arr {
		fmt.Print((blanks[0 : (array_len*2-index)/2]))
		for _, val := range value {
			fmt.Print(strconv.Itoa(val) + " ")
		}
		fmt.Println()
	}
}

func generate(numrows int) [][]int {
	pascal := [][]int{}
	for i := 0; i < numrows; i++ {
		sl := make([]int, i+1, i+1)
		sl[0] = 1
		sl[i] = 1
		for j := 0; j <= i; j++ {
			if j > 0 && j < i {
				sl[j] = pascal[i-1][j-1] + pascal[i-1][j]
			}
		}
		pascal = append(pascal, sl)
	}
	return pascal
}
