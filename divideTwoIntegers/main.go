package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x := 2
	y := 4
	fmt.Println(x << 1)
	fmt.Println(y >> 1)
	fmt.Println(strconv.Itoa(int(divide(100, 3))))
}

func divide(divided int32, divisor int32) int32 {
	m := int32(math.Abs(float64(divided)))
	n := int32(math.Abs(float64(divisor)))
	res := 0
	if n == 1 {
		return divided
	}
	for m >= n {
		t := n
		p := 1
		for m >= (t << 1) {
			p <<= 1
			t <<= 1
		}
		res += p
		m -= t
	}
	if (divided < 0 && divisor > 0) || (divided > 0 && divisor < 0) {
		return -int32(res)
	} else {
		return int32(res)
	}
}
