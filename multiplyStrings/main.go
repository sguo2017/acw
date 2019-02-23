package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var a = '2'
	var b = '0'
	var n = a - b
	fmt.Println(strconv.Itoa(int(n)))
	fmt.Printf(solution("457", "26"))
}

func solution(num1 string, num2 string) string {
	var (
		length = len(num1) + len(num2)
		num    = []int{}
		b1     = []byte(num1)
		b2     = []byte(num2)
		i      = len(b1) - 1
		j      = len(b2) - 1
		//buf  string
		buffer bytes.Buffer
	)
	num = make([]int, length)
	for ; i >= 0; i-- {
		for j = len(b2) - 1; j >= 0; j-- {
			p1 := i + j
			p2 := i + j + 1
			sum := num[p2] + int(b1[i]-'0')*int(b2[j]-'0')
			num[p1] += sum / 10
			num[p2] = sum % 10
		}
	}
	for i = 0; i <= length-1; i++ {
		if !(num[i] == 0 && buffer.Len() == 0) {
			buffer.WriteString(strconv.Itoa(num[i]))
			//buf = buf + strconv.Itoa(num[i])
		}
	}
	return buffer.String()
}
