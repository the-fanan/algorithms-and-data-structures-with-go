package main


import (
	"fmt"
	"errors"
)

func main() {
	d := &DoublyLinkedList{}
	d.Push(5)
	d.Push(10)
	d.Push(15)
	d.Push(20)
	d.Print()
	d.Pop()//20
	d.Print()
	d.Shift()//5
	d.Print()
	d.Unshift(5)//5
	d.Print()
	n,e := d.Get(1)
	fmt.Println(n,e)
	d.Set(1,8)
	d.Print()
	d.Insert(0,2)
	d.Print()
	d.Insert(1,6)
	d.Print()
	d.Insert(3,10)
	d.Print()
	d.Insert(6,20)
	d.Print()
	d.Remove(4)
	d.Print()
	d.Reverse()
	d.Print()
}

type Node struct {
	Next *Node
	Previous *Node
	Value int
}

/**
	* Like singly list list but traverseing is easier due to extran pointer
	* Drawback is that it uses more memory
	*/
type DoublyLinkedList struct {
	Length int 
	Head *Node 
	Tail *Node
}

func (d *DoublyLinkedList) Print(){
	if d.Head == nil{
		fmt.Print("List is empty")
	} else {
		current := d.Head
		i := 0
		for current != nil {
			fmt.Print(i, ": ", current.Value, ", ")
			current = current.Next
			i++
		}
	}
	fmt.Println("")
	fmt.Println("Length:", d.Length)
}

func (d *DoublyLinkedList) Reverse() {
	newTail := d.Head 
	newHead := &Node{}
	current := d.Head 
	for current != nil {
		previous := current.Previous 
		next := current.Next 
		current.Previous = next 
		current.Next = previous 
		current = next 
		if current != nil {
			newHead = current
		}
	}
	d.Tail = newTail 
	d.Head = newHead
}

func (d *DoublyLinkedList) Remove(k int) error {
	if k < 0 || k > d.Length - 1 {
		return errors.New("Index out of range")
	}

	if k == 0 {
		d.Shift()
		return nil
	}

	if k == d.Length - 1 {
		d.Pop() 
		return nil
	}

	node,_ := d.Get(k)
	p := node.Previous
	n := node.Next 
	p.Next = n 
	n.Previous = p 
	node.Next = nil 
	node.Previous = nil

	d.Length--
	return nil
}

func (d *DoublyLinkedList) Insert(k,v int) error {
	if k < 0 || k > d.Length {
		return errors.New("Index out of range")
	}

	n := &Node{Value: v}
	if d.Length == 0 {
		d.Head = n 
		d.Tail = n
		return nil
	}

	if d.Length == k {
		d.Push(v)
		return nil
	}

	if k == 0 {
		d.Unshift(v)
		return nil
	}

	useHead := true 
	if d.Length - k < k {
		useHead = false 
	}
	current := &Node{}
	if useHead {
		current = d.Head 
		for i:= 0; i < k; i++ {
			current = current.Next
		}
	} else {
		current = d.Tail 
		for i := d.Length - 1; i > k; i--{
			current = current.Previous
		}
	}
	previous := current.Previous
	previous.Next = n 
	current.Previous = n 
	n.Next = current 
	n.Previous = previous
	d.Length++
	return nil
}

func (d *DoublyLinkedList) Set(k,v int) {
	n, _ := d.Get(k)
	n.Value = v
}

func (d *DoublyLinkedList) Get(k int) (*Node,error) {
	if d.Length == 0 {
		return nil, errors.New("List is empty")
	}
	if k < 0 || k > d.Length - 1 {
		return nil, errors.New("Index out of range")
	}
	useHead := true 
	if d.Length - k < k {
		useHead = false
	}

	current := &Node{}
	if useHead {
		current = d.Head 
		for i := 0; i < k; i++ {
			current = current.Next
		}
	} else {
		current = d.Tail 
		for i := d.Length - 1; i > k; i-- {
			current = current.Previous 
		}
	}
	return current,nil
}

func (d *DoublyLinkedList) Unshift(v int) {
	n := &Node{Value: v}
	if d.Length == 0 {
		d.Tail = n
	} else {
		n.Next = d.Head 
		d.Head.Previous = n 
	}
	d.Head = n
	d.Length++
}

//remove from begining
func (d *DoublyLinkedList) Shift() *Node{
	if d.Length == 0 {
		return nil
	}
	n := d.Head 
	if d.Length == 1 {
		d.Head = nil 
		d.Tail = nil 
	} else {
		d.Head = d.Head.Next 
		d.Head.Previous = nil 
	}
	n.Next = nil 
	d.Length--
	return n
}

func (d *DoublyLinkedList) Pop() *Node{
	if d.Length == 0 {
		return nil
	}
	n := d.Tail 

	if d.Length == 1 {
		d.Head = nil 
		d.Tail = nil 
	} else {
		d.Tail = d.Tail.Previous
		d.Tail.Next = nil
	}
	n.Previous = nil 
	d.Length--
	return n
}

func (d *DoublyLinkedList) Push(v int) {
	n := &Node{Value: v}
	if d.Head == nil {
		d.Head = n 
	} else {
		d.Tail.Next = n 
		n.Previous = d.Tail 	
	}
	d.Tail = n
	d.Length++
}
