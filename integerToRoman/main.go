package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intToRoman(23))
	fmt.Println(intToRoman2(23))
}

func intToRoman(num int) string {
	bufStr := bytes.NewBufferString("")
	nums := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNums := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < 13; i++ {
		tmp := num / nums[i]
		if tmp > 0 {
			bufStr.Write(bytes.Repeat([]byte(romanNums[i]), tmp))
			num -= tmp * nums[i]
		} else {
			continue
		}
	}
	return bufStr.String()
}

var itomanMap = map[int]string{
	1:    "I",
	2:    "II",
	3:    "III",
	4:    "IV",
	5:    "V",
	6:    "VI",
	7:    "VII",
	8:    "VIII",
	9:    "IX",
	10:   "X",
	20:   "XX",
	30:   "XXX",
	40:   "XL",
	50:   "L",
	60:   "LX",
	70:   "LXX",
	80:   "LXXX",
	90:   "XC",
	100:  "C",
	200:  "CC",
	300:  "CCC",
	400:  "CD",
	500:  "D",
	600:  "DC",
	700:  "DCC",
	800:  "DCCC",
	900:  "CM",
	1000: "M",
	2000: "MM",
	3000: "MMM",
}

func intToRoman2(num int) string {
	result := ""
	flag := 1
	for num != 0 {
		i := num % 10
		result = itomanMap[i*flag] + result
		num, flag = num/10, flag*10
	}
	return result
}
