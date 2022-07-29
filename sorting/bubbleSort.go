package main 

import (
	"fmt"
)

func main(){
    fmt.Println("Bubble Sort:")
  	fmt.Println(bubbleSort([]int{5,2,3,22,-4,0,11,-5,1,1}))
    fmt.Println(bubbleSort([]int{-5,12,3,2,-4,0,11,-5,1,1}))
    fmt.Println(bubbleSort([]int{-5,-2,3,112,-4,0,100,-5,1,2}))
}

func bubbleSort(arr []int ) []int {
	for i:= 0 ; i < len(arr) - 1 ; i++ {
		for j:=0 ; j < len(arr) - 1 - i ; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j+1] = arr[j+1],arr[j]
			} 		
		} 
	}
	return arr
}
