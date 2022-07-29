package main
import "fmt"


type node struct {
  data int 
  next *node
}

type queue struct {
  head *node
  tail *node
  size int
}


func main(){
  q:= queue{nil,nil,10}
  fmt.Println(q)

  q.enqueue(10)
  q.enqueue(20)
  q.enqueue(0)

  fmt.Println(*q.find(10))

  q.display()
} 

func (q *queue) enqueue(data int){
  n:= node{data, nil}
  if q.head == nil {
    q.head = &n
    q.tail = &n
  }else {  
    q.tail.next = &n
    q.tail = q.tail.next
  }
  q.size++
}


func (q *queue) dequeue () int {
  data := q.head.data
  q.head = q.head.next
  q.size--
  if q.size == 0 {
    q.tail = nil
  }
  return data
}

func (q *queue) find(data int) *node{
  tmp := q.head

  for tmp != nil {
    if tmp.data == data {
      return tmp
    }
    tmp = tmp.next
  }
  return nil
}

func (q *queue) display(){
  tmp := q.head
  for tmp.next != nil {
    fmt.Print(tmp.data ," --> ")
    tmp = tmp.next
  }
  if &tmp != nil {
    fmt.Print(tmp.data )
  }
}