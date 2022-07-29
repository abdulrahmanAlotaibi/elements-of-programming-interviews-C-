package main 

import (
	"fmt"
	"errors"
	"log"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
	
}

type Stack [] string

func main (){

	var stack Stack

	stack.Push("John Wick")
	stack.Push("Spiderman")
	stack.Push("Mary Jane")
	stack.Push("Superman")


	stack.Display()

  item ,err:=stack.peek()
  Check(err)
  fmt.Println("Peek Item value: ", item)
  
	item1, err :=stack.Pop()
	Check(err)

	fmt.Println("Item1 value: ", item1)

	stack.Display()

	item2, err :=stack.Pop()
	Check(err)

	fmt.Println("Item2 value: ", item2)

	item3,err :=stack.Pop()

	Check(err)
	fmt.Println("Item3 value: ", item3)


	// Empty stack : This will throw an error
	item4,err :=stack.Pop() 
	Check(err)
	
	fmt.Println("Item4 value: ", item4)

	// Display the stack items
	stack.Display()

}

func (stack *Stack) Pop () (string, error){
	 
  stack.isEmpty()	

	item := (*stack)[len(*stack) - 1]

	*stack = (*stack)[:len(*stack) - 1]
	fmt.Println("Item has been popped out successfully")

	return item, nil
}

func (stack *Stack) Push(item string) {

	*stack = append(*stack, item)

	fmt.Println("Item has been pushed successfully")
}

func (stack *Stack) Display(){

	for i, item := range *stack {
		fmt.Printf("[%d] : %s\n", i, item)
	}
	
}

func (stack *Stack) isEmpty() {
  if len(*stack) == 0 {
		errors.New("The stack is empty")
	}  
}

func (stack *Stack) peek()(string, error) {
  stack.isEmpty()
	item := (*stack)[len(*stack) - 1]
  return item, nil
} 

