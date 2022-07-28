package main

import "fmt"

const  (
  PreOrder  = "preOrder"
  InOrder  = "inOrder"
  PostOrder = "postOrder"
  Root = "root"
  LeftChild = "leftChild"
  RightChild = "rightChild"
  Parent = "parent"
)

type node struct {
  data int 
  left *node
  right *node
}

type binaryTree struct {
  root *node 
  current *node
}

func main(){
    var bt binaryTree = binaryTree{nil, nil}
    
    bt.insert(90, Root)
    fmt.Println(bt.current)
    bt.insert(20, RightChild)
    bt.find(Parent)
    bt.insert(16, LeftChild)
    bt.find(Parent)
  
    fmt.Println(bt.current)

    bt.find(LeftChild)
    
    bt.deleteSubTree()

    fmt.Println(bt.current)
}

func (bt *binaryTree) insert(data int, relative string ) bool {
  switch relative {
    case Root:
      if !bt.isTreeEmpty() {
        return false
      }
      n := node{data,nil,nil}
      bt.root = &n
      bt.current = bt.root
      return true
    case Parent:{
      // You can't find the parent of the root
      if bt.current == bt.root {
        return false 
      }
      bt.current = findParent(bt.current, bt.root)
      return true
    }
    case LeftChild:
      if bt.current.left != nil {
        return false
      }
      n := node{data, nil, nil}
      bt.current.left = &n
      bt.current = bt.current.left
      return true
    case RightChild:
      if bt.current.right != nil {
        return false
      }
      n := node{data, nil, nil}
      bt.current.right = &n
      bt.current = bt.current.right
      return true
    default:
      return false
  }
}

func (bt *binaryTree) deleteSubTree() {
  if bt.current == bt.root {
    bt.root = nil
    bt.current = bt.root
  } else {
    n := bt.current
    bt.find(Parent)
    if(bt.current.left == n){
      bt.current.left = nil
    } else if (bt.current.right == n){
      bt.current.right = nil
    }
    bt.current = bt.root
  }
  
}

func (bt *binaryTree) update(data int){
  bt.current.data = data
}

func (bt *binaryTree) retrieve(data int) *node{
  return bt.current
}

func  (bt *binaryTree) find(relative string)bool{
  switch relative{
    case Root :{
      bt.current = bt.root
      return true
    }
    case Parent:{
      if bt.current == bt.root {
        return false
      }
      bt.current = findParent(bt.current, bt.root)
      return true
    }
    case LeftChild:{
      if bt.current.left == nil {
        return false
      }
      bt.current  = bt.current.left
      return true
    }
    case RightChild:{
      if bt.current.right == nil {
        return false
      }
      bt.current  = bt.current.right
      return true
    }
    default:{
      return false
    }
  }

}

func findParent (n *node, root *node)*node{
  if root == nil {
    return nil
  }
  if root.right == nil && root.left == nil {
    return nil
  } else if root.right == n || root.left == n {
      return root
  }else {
    tmp:= findParent(n, root.right)
    if tmp != nil {
      return tmp.left
    }else {
      return findParent(n, root.left)
    }
  }
} 

func (bt *binaryTree) isTreeEmpty() bool {
  return bt.root == nil
}
