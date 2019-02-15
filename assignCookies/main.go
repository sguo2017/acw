package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findContentChildren(chidren []int, cookies []int) int {
	var (
		result         = 0
		i              = 0
		j              = 0
		lengthChildren = len(chidren)
		lengthCookies  = len(cookies)
	)
	sort.Ints(chidren)
	sort.Ints(cookies)
	for ; i < lengthChildren && j < lengthCookies; i++ {
		if cookies[j] >= chidren[i] {
			result++
			j++
		}
	}
	return result
}

func main() {
	var children = []int{1, 4, 2, 3}
	var cookies = []int{2, 1, 2}
	fmt.Printf(strconv.Itoa(findContentChildren(children, cookies)))

}
