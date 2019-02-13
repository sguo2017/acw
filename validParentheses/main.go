package main

import (
	"container/list"
	"fmt"
)


var myMap = map[string]int32{}


func main() {
	myMap["}"]=123
	myMap["]"]=91
	myMap[")"]=40
	fmt.Println(solution("{}"))
	fmt.Println(solution("{}}"))
	fmt.Println(solution("{[]}"))
	fmt.Println(solution("{{}"))
}

func solution(s string) bool {

	stack := list.New()

	for _, ch := range s  {
		index := string(ch)
		// 如果还有右侧
		if(myMap[index] > 0){
			// 没有多余 false
			if(stack.Len() == 0){
				return false
			}else{
				// 弹出一个右侧 不匹配 false
				backElement := stack.Back()
				topElement := backElement.Value.(int32);
				stack.Remove(backElement)
				if(topElement != myMap[index]){
					return false
				}
			}


		}else {
			stack.PushBack(ch)
		}
	}

	return stack.Len()==0

}
