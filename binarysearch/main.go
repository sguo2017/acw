package main

import "fmt"

func main() {
	var arr = []int32{11,12,13,14,15,16,17}
	fmt.Print("test\n")
	fmt.Printf("seach arr for 13 index is %d\n", binarysearch(arr, 13))
}

func binarysearch(nums []int32, target int32) int{
	low := 0
	high := len(nums)
	for{
		if(high < low){
			break
		}

		middle := (low+high)/2
		if(target == nums[middle-1]){
			return middle;
		}else if(target > nums[middle-1]){
			low = middle+1
		}else {
			high = middle-1;
		}
	}

	return -1

}