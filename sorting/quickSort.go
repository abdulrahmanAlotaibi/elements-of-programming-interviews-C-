package main 

import (
	"fmt"
)

func main(){
  	fmt.Println(quickSort([]int{2,1,-3,0}, 0 , 3))
}

func quickSort(arr[]int, low int, high int) []int{
  if low < high {
    pivot := partition(arr, low , high)
    quickSort(arr, low, pivot - 1 )
    quickSort(arr, pivot + 1, high)
  }

  return arr
}

func partition(arr[]int, low int ,  high int) int{
  pivot := arr[high]
  i := low

  for j:= i ; j < high ; j++ {
    if arr[j] < pivot {
      arr[j], arr[i] = arr[i], arr[j]
      i++
    }
  }

  arr[high], arr[i] = arr[i] , arr[high] 

  return i
}