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

func KMP(content []byte, sub []byte) (index int) {
	var (
		next    []int = getNext(sub)
		lenghtC int   = len(content)
		lenghtS int   = len(sub)
		i       int   = 0
		j       int   = 0
	)
	for i < lenghtC && j < lenghtS {
		if j == -1 || content[i] == sub[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == lenghtS {
		return i - j
	} else {
		return -1
	}

}

func main() {
	content := []byte("why every programming language use the hello world as the first test???")
	sub := []byte("hello world")
	fmt.Println(KMP(content, sub))
}
