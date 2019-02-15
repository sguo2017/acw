package main

import (
	"fmt"
	"math"
	"strconv"
)

func base7(num int) string {
	if num == 0 {
		return "0"
	}
	var (
		str = ""
		tmp = int(math.Abs(float64(num)))
		bit = 0
	)
	for tmp != 0 {
		bit = tmp % 7
		str = strconv.Itoa(bit) + str
		tmp = tmp / 7
	}
	if num < 0 {
		str = "-" + str
	}
	return str
}

func main() {
	fmt.Printf(base7(14))
}
