package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var arr = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("result:" + strconv.Itoa(trap(arr)))
}

func trap(height []int) int {
	len := len(height)
	left := make([]int, len, len)
	right := make([]int, len, len)
	var max int = 0
	for i := 0; i < len; i++ {
		if height[i] > max {
			max = height[i]
		}
		left[i] = max
	}
	max = 0
	for j := len - 1; j >= 0; j-- {
		if height[j] > max {
			max = height[j]
		}
		right[j] = max
	}
	var sum int = 0
	for i := 1; i < len-1; i++ {
		m := int(math.Min(float64(left[i-1]), float64(right[i+1]))) - height[i]
		if m > 0 {
			sum += m
		}
	}
	return sum
}
