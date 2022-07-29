package main 

import (
	"fmt"
)

func main(){
  fmt.Println("Merge Sort:")
  fmt.Println(quickSort([]int{22,11,-3,0}, 0 , 3))
  fmt.Println(quickSort([]int{9,1,-3,0}, 0 , 3))
  fmt.Println(quickSort([]int{-2,11,-3,10,0,220}, 0 , 3))
}

func quickSort(arr[]int, low int, high int) []int {
  if low < high {
    pivot := partition(arr, low , high)
    quickSort(arr, low, pivot - 1 )
    quickSort(arr, pivot + 1, high)
  }

  return arr
}

func partition(arr[]int, low int ,  high int) int {
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