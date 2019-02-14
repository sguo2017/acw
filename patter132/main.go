package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	var arr1 = []int32{1, 2, 3, 4}
	var arr2 = []int32{3, 1, 4, 2}
	var arr3 = []int32{-1, 3, 2, 0}

	fmt.Println("solution1")
	fmt.Println("arr1:" + strconv.FormatBool(solution1(arr1)))
	fmt.Println("arr2:" + strconv.FormatBool(solution1(arr2)))
	fmt.Println("arr3:" + strconv.FormatBool(solution1(arr3)))

	fmt.Println("\n\nsolution2")
	fmt.Println("arr1:" + strconv.FormatBool(solution2(arr1)))
	fmt.Println("arr2:" + strconv.FormatBool(solution2(arr2)))
	fmt.Println("arr3:" + strconv.FormatBool(solution2(arr3)))

	fmt.Println("\n\nsolution3")
	fmt.Println("arr1:" + strconv.FormatBool(solution2(arr1)))
	fmt.Println("arr2:" + strconv.FormatBool(solution2(arr2)))
	fmt.Println("arr3:" + strconv.FormatBool(solution2(arr3)))
}

const INT_MAX = int(^uint(0) >> 2)
const INT_MIN = ^INT_MAX

func solution1(arr []int32) bool {
	var i, j int
	temp := INT_MAX
	n := len(arr)
	for i = 0; i < n-1; i++ {
		temp = Min(temp, int(arr[i]))
		if temp == int(arr[i]) {
			continue
		}
		for j = n - 1; j > i; j-- {
			if int(arr[j]) > temp && arr[j] < arr[i] {
				return true
			}
		}
	}
	return false
}

func solution2(arr []int32) bool {
	var i, j, k int
	n := len(arr)
	for i = 0; i < n-1; {
		if arr[i] > arr[i+1] {
			i++
			continue
		}
		for j = i + 1; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				break
			}
		}
		for k = j + 1; k < n-1; k++ {
			if arr[k] > arr[i] && arr[k] < arr[i] {
				return true
			}
		}
		i = j + 1
	}
	return false
}

func solution3(arr []int32) bool {
	var stack = list.New()
	third := INT_MIN
	var i int
	length := len(arr)
	for i = length - 1; i > 0; i-- {
		if int(arr[i]) < third {
			return true
		}
		for stack.Len() > 0 {
			backElement := stack.Back()
			if backElement.Value.(int32) > arr[i] {
				third = int(backElement.Value.(int32))
				stack.Remove(backElement)
			}
		}
		stack.PushBack(arr[i])
	}
	return false
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
