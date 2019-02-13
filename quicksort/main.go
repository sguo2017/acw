package main

import "fmt"

func main() {
	fmt.Print("quickSort")

	var arr = []int32{19,2,3,9,18,7,10,16}

	quickSort(arr, 0, 7)

}

func quickSort(arr []int32, low int32, high int32)  {
	l := low
	h := high
	povit := arr[low]
	fmt.Printf("l=%d, h=%d,povit=%d \n", l+1, h+1, povit)
	for {
		if(l > h){
			break
		}

		for{
			if(l > h || arr[h] <= povit){
				break
			}

			h--

			if(l < h){
				arr[l],arr[h] = arr[h],arr[l]
				l++
			}

		}

		for{
			if(l>h || arr[l] >= povit){
				break
			}

			l++

			if(l<h){
				arr[l],arr[h] = arr[h],arr[l]
				h--
			}
		}


		if(l>low){
			quickSort(arr, low, l-1)
		}

		if(h<high){
			quickSort(arr, l+1, high)
		}

	}
}