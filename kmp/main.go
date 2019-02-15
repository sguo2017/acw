package main

// https://blog.csdn.net/x__1998/article/details/79951598

import (
	"fmt"
)

func getNext(sub []byte) (next []int) {
	var (
		length = len(sub)
		i      int
		j      int
	)
	next = make([]int, length)
	i = 0
	j = -1
	next[0] = -1
	for i < length-1 {
		if j == -1 || sub[i] == sub[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func KMP(content []byte, start_index int, end_index int, sub []byte) (index int) {
	var (
		next       []int = getNext(sub)
		sub_index  int   = 0
		sub_length int   = len(sub)
	)
	for i := start_index; i <= end_index; i++ {
		if content[i] == sub[sub_index] {
			match_start := i
			for j := sub_index; j <= sub_length; j++ {
				if j == sub_length {
					return match_start - sub_index
				}
				if i >= end_index || content[i] != sub[j] {
					sub_index = next[j]
					break
				}
				i++
			}
		}
	}

	return -1
}

func main() {
	content := []byte("why every programming language use the hello world as the first test???")
	sub := []byte("hello world")
	fmt.Println(KMP(content, 0, len(content)-1, sub))
}
