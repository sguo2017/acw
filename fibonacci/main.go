package main

import "fmt"

func main() {
	fmt.Printf("test\n")
	fmt.Printf("fibonacci 5 return %d\n", calc(5))
	fmt.Printf("fibonacci 6 return %d\n", calc(6))
}


func calc(n int32) int32 {
	if(n==0 || n==1){
		return 1
	}

	return calc(n-1)+calc(n-2)
}
