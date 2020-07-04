package main 

import (
	"fmt"
	"errors"
)

func main(){
	s := &Stack{}
	s.Put(5)
	s.Put(10)
	s.Put(15)
	s.Put(20)
	s.Print()
	val, err := s.Pop()
	fmt.Println(val, err)
	s.Print()
	s.Pop()//15
}

type Node struct {
	Next *Node 
	Value int
}

/**
	* Follows Last In First Out principle (LIFO)
	*/
type Stack struct {
	Head *Node 
	Length int
}

func (s *Stack) Print(){
	node := s.Head 
	for node != nil {
		fmt.Print(node.Value, " ")
		node = node.Next
		if node == nil {
			fmt.Println("")
		}
	}
	fmt.Println("Length:", s.Length)
}

func (s *Stack) Put(v int) error {
	if s.Length > 10000 {
		return errors.New("Stack Overflow")
	}
	n := &Node{Value: v}
	if s.Length == 0 {
		s.Head = n 
	} else {
		n.Next = s.Head 
		s.Head = n
	}
	s.Length++
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.Length == 0 {
		return 0, errors.New("The stack is empty")
	}
	node := s.Head
	s.Head = s.Head.Next 
	node.Next = nil 
	s.Length--
	return node.Value,nil
}