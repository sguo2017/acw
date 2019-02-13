package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("hello \n")
	fmt.Printf(say("11211122"))
}

func say(s string) string {
	var result string
	var temp int32
	count := 0
	for _, ch := range s {
		if ch == temp {
			count++
		} else {
			if count > 0 {
				result += strconv.Itoa(count) + string(temp)
			}
			temp = ch
			count = 1
		}
	}
	if count > 0 {
		result += strconv.Itoa(count) + string(temp)
	}
	return result
}
