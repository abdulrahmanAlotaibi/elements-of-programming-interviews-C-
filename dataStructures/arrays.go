package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sort even numbers first then odd numbers:")
  evenOdd([]int{3,1,2,4,5})
  evenOdd([]int{4,4,2,13,9,5})
  evenOdd([]int{13,9,5,4,4,2,5})

  fmt.Println("Dutch National Flag Algorithm:")
  fmt.Println(dutchFlagPartition([]int{13,9,5,4,4,2,51}, 2))
  fmt.Println("Improved Dutch National Flag:")
  fmt.Println(improvedDutchFlag([]int{13,9,5,4,4,2,51}, 2))
  fmt.Println("Delete any duplicates in an array")
  fmt.Println(deleteDuplicatesSortedArray([]int{2,3,5,7,11,11,13}))
}

func evenOdd(nums []int){
  for i, j := 0 , len(nums) - 1 ; i <= j;{
        if(nums[i] % 2 != 0 && nums[j] % 2 == 0){ // Even
          nums[i], nums[j] = nums[j],nums[i]
        }  
        if nums[i] % 2 == 0 {
          i++
        }
        if nums[j] % 2 != 0 {
          j--
        }
  }

  fmt.Println(nums)
}


func dutchFlagPartition(arr []int, pivotIndex int) []int{
  pivot := arr[pivotIndex]
  for i:= 0 ; i < len(arr) ; i++ {

    for j:= i + 1 ; j < len(arr) ; j++ {
      if arr[j] < pivot {
        arr[i], arr[j] = arr[j], arr[i]
        break;
      }
    }
    
  }
  // Second pass 
  for i:= len(arr) - 1 ; i >= 0 ; i-- {
    for j:= i - 1 ; j >= 0 ; j-- {
      if arr[j] > pivot {
        arr[i], arr[j] = arr[j], arr[i]
        break;
      }
    }
  }

  return arr
}

func improvedDutchFlag(arr []int, pivotIndex int)[]int{
  pivot := arr[pivotIndex]
  smaller := 0;

  for i:=0; i < len(arr) ; i++ {
    if arr[i] < pivot {
      arr[i], arr[smaller] = arr[smaller] , arr[i]
      smaller++
    }
  }

  larger := len(arr) - 1 
  for i:= len(arr) - 1 ; i >= 0 ; i-- {
    if arr[i] > pivot {
      arr[i], arr[larger] = arr[larger] , arr[i]
      larger--
    }
  }


  return arr
}

func deleteDuplicatesSortedArray(arr []int)[]int{
  for i:= 0 ; i < len(arr) - 1 ; i++ {
    if arr[i] == arr[i+1]{
      arr = append(arr[:i], arr[i+1:]...)
    }
  }
  return arr
}