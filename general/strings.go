package main 
import "fmt"

func main(){
  fmt.Println("Is the the string the same if we reverse it ?")
  fmt.Println(isPalindromic("AHA"))
  fmt.Println(isPalindromic("Go"))
}

func isPalindromic(s string)bool{
  for i,j := 0, len(s) - 1 ; i < len(s) ; i,j = i+1, j-1 {
    if s[i] != s[j]{
      return false
    }
  }
  return true
}

