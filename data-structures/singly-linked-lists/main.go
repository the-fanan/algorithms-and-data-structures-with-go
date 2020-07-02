package main

import (
	"fmt"
	"errors"
)

func main(){
	s := &SinglyLinkedList{} 
	s.Push(5)
	s.Push(10)
	s.Push(15)
	s.Print()

	s.Pop()
	s.Print()

	s.Pop()
	s.Print()

	s.Shift()
	s.Print()

	s.Push(15)
	s.Print()

	s.Unshift(20)
	s.Print()
	
	v,e := s.Get(10)
	fmt.Println(v,e)
	v,e = s.Get(1)
	fmt.Println(v,e)
	v,e = s.Set(0,5)
	fmt.Println(v,e)

	s.Insert(1, 10)
	s.Print()

	s.Insert(1, 8)
	s.Print()

	s.Insert(4, 20)
	s.Print()

	s.Remove(2)
	s.Print()

	s.Insert(4, 25)
	s.Print()

	s.Reverse2()
	s.Print()
}


type Node struct {
	Value int
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node 
	Tail *Node
	Length int
}

func (s *SinglyLinkedList) Print(){
	if s.Head == nil{
		fmt.Println("List is empty")
	} else {
		current := s.Head
		i := 0
		for current != nil {
			fmt.Print(i, ": ", current.Value, ", ")
			current = current.Next
			i++
		}
	}
	fmt.Println("")
}

func (s *SinglyLinkedList) Reverse() {
	n := s.Length
	for i := n - 1; i > (n/2) - 1; i--{
		l,_ := s.Get(n - i - 1)
		r,_ := s.Get(i)
		left := l.Value 
		right := r.Value
		s.Set(i, left)
		s.Set(n - i - 1, right)
	}
}

func (s *SinglyLinkedList) Reverse2() {
	oldNext := &Node{} 
	oldNext = nil
	current := s.Head
	previous := &Node{}
	previous = nil
	newTail := s.Head 
	newHead := &Node{}

	for current != nil{
		oldNext = current.Next 
		current.Next = previous 
		previous = current 
		current = oldNext
		if current != nil {
			newHead = current
		}
	}
	s.Head = newHead 
	s.Tail = newTail
}

func (s *SinglyLinkedList) Remove(p int) (*Node, error) {
	if p < 0 || p >= s.Length {
		err := errors.New("Index is out of range")
		return nil,err
	} 

	if p == 0 {
		//remove the head
		return s.Shift(), nil
	}

	prev := &Node{}
	prev = nil 
	current := s.Head
	for i := 0; i < p; i++ {
		prev = current 
		current = current.Next
	}
	prev.Next = current.Next
	s.Length--
	return current,nil
}

func (s *SinglyLinkedList) Insert(position int, value int) (*Node, error) {
	if position < 0 || position > s.Length {
		err := errors.New("Index is out of range")
		return nil,err
	}

	i := 1
	o := s.Head
	n := &Node{Value: value}
	for i < position {
		o = o.Next
		i++
	}
	old := o.Next 
	o.Next = n 
	n.Next = old 
	s.Length++
	return n,nil
}

func (s *SinglyLinkedList) Set(position int, value int) (*Node, error) {
	if position < 0 || position > s.Length - 1 {
		err := errors.New("Index is out of range")
		return nil,err
	}
	i := 0
	n := s.Head
	for i < position {
		n = n.Next
		i++
	}
	n.Value = value
	return n,nil
}

func (s *SinglyLinkedList) Get(p int) (*Node, error) {
	if p < 0 || p > s.Length - 1 {
		err := errors.New("Index is out of range")
		return nil,err
	}
	i := 0
	n := s.Head
	for i < p {
		n = n.Next
		i++
	}
	return n,nil
}

func (s *SinglyLinkedList) Unshift(v int) {
	n := &Node{Value: v}
	if s.Head == nil {
		s.Head = n
		s.Tail = n
	} else {
		n.Next  = s.Head 
		s.Head = n
	}
	s.Length++
}

func (s *SinglyLinkedList) Shift() *Node{
	oldHead := s.Head
	if s.Head == s.Tail {
		s.Head = nil
		s.Tail = nil
	} else {
		s.Head = s.Head.Next
	}
	s.Length--
	return oldHead
}

func (s *SinglyLinkedList) Pop() *Node{
	prev := &Node{} 
	prev = nil
	current := s.Head
	if s.Head == s.Tail {
		s.Head = nil 
		s.Tail = nil 
		s.Length = 0
	} else {
		for current.Next != nil  {
			prev = current 
			current = current.Next
		}
		if prev != nil {
			prev.Next = nil 
			s.Tail = prev
		}
		s.Length--
	}
	return current
}

func (s *SinglyLinkedList) Push(v int) *Node{
	n := &Node{Value: v}
	if s.Head == nil {
		s.Head = n
	} else {
		s.Tail.Next = n
	}
	s.Tail = n
	s.Length++
	return n
}