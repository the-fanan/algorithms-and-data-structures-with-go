package main

import (
	"fmt"
	"errors"
)

type QNode struct {
	Next *QNode 
	Value *Node
}

/**
	* For tree traversing
	*/
type Queue struct {
	Head *QNode 
	Tail *QNode
	Length int
}

func (q *Queue) Print(){
	node := q.Head 
	for node != nil {
		fmt.Print(node.Value, " ")
		node = node.Next
		if node == nil {
			fmt.Println("")
		}
	}
	fmt.Println("Length:", q.Length)
}

func (q *Queue) Dequeue() (*Node,error) {
	if q.Length == 0 {
		return nil, errors.New("Queue is empty")
	}
	node := q.Head 
	q.Head = q.Head.Next 
  node.Next = nil
	q.Length--
	return node.Value, nil
}

func (q *Queue) Enqueue(v *Node) {
	n := &QNode{Value: v}
	if q.Length == 0 {
		q.Head = n
	} else {
		q.Tail.Next = n
	}
	q.Tail = n
	q.Length++
}

type SNode struct {
	Next *SNode 
	Value *Node
}

/**
	* Follows Last In First Out principle (LIFO)
	*/
type Stack struct {
	Head *SNode 
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

func (s *Stack) Put(v *Node) error {
	if s.Length > 10000 {
		return errors.New("Stack Overflow")
	}
	n := &SNode{Value: v}
	if s.Length == 0 {
		s.Head = n 
	} else {
		n.Next = s.Head 
		s.Head = n
	}
	s.Length++
	return nil
}

func (s *Stack) Pop() (*Node, error) {
	if s.Length == 0 {
		return nil, errors.New("Stack is empty")
	}
	node := s.Head
	s.Head = s.Head.Next 
	node.Next = nil 
	s.Length--
	return node.Value,nil
}

func main(){
	b := &BST{}
	b.Insert(10)//root
	b.Insert(13)
	b.Insert(5)
	b.Insert(16)
	b.Insert(12)
	b.Insert(6)
	b.Insert(4)
	fmt.Println(b.Find(6))
	fmt.Println(b.Find(26))
	b.BFS()
	b.DFSPreO()
	b.DFSPostO()
	b.DFSInO()
}

/**
	* For a BST tree, items to the left of the node must be smaller than the node's value. Items to the right must be greater than the node's value.
	*/
type Node struct {
	Left *Node 
	Right *Node 
	Value int
	Count int
}

type BST struct {
	Root *Node
}

func (b *BST) Insert(v int) {
	node := &Node{Value: v}
	node.Count++
	if b.Root == nil {
		b.Root = node 
	} else {
		current := b.Root
		for current != nil {
			if node.Value == current.Value {
				current.Count++ 
				break
			}
			if node.Value > current.Value {
				//do I have a right? 
				if current.Right == nil {
					current.Right = node 
					break
				} else {
					current = current.Right
				}
			} else if node.Value < current.Value {
				if current.Left == nil {
					current.Left = node 
					break 
				} else {
					current = current.Left
				}
			}
		}
	}
}

func (b *BST) Find(v int) bool{
	if b.Root == nil {
		return false 
	} else {
		current := b.Root
		for current != nil {
			if v == current.Value {
				return true
			}
			if v > current.Value {
				if current.Right == nil {
					return false
				} else {
					current = current.Right
				}
			} else if v < current.Value {
				if current.Left == nil {
					return false
				} else {
					current = current.Left
				}
			}
		}
	}
	return false
}
/**
	* Traversing Trees (touching each node only once)
	* 1. Breadth First search
	* 2. Depth First Pre-Order
	* 3. Depth First Post-Order 
	* 4. Depth First In-Order
*/
//Go through siblings before you go to their children
func (b *BST) BFS(){
	if b.Root == nil {
		return
	} else {
		q := &Queue{} 
		q.Enqueue(b.Root)
		i := 1
		br := 1
		for q.Length > 0 {
			node, err := q.Dequeue() 
			if err == nil {
				//the queue was not empty
				fmt.Print(node.Value, " ")
				if node.Left != nil {
					q.Enqueue(node.Left)
				}
				if node.Right != nil {
					q.Enqueue(node.Right)
				}
			} 
			if i == br {
				fmt.Println("")
				i = 1 
				br = br * 2
			} else {
				i++
			}
		}
	}
	fmt.Println("")
}

//explore current node, then look at left, then right -- to export or duplicate a tree so it's structure is the same when used to reconstruct a new tree
func (b *BST) DFSPreO(){
	b.DFSPreOHelper(b.Root)
	fmt.Println("")
}

func (b *BST) DFSPreOHelper(n *Node){
	if n == nil {
		return
	} else {
		fmt.Print(n.Value, " ")
		b.DFSPreOHelper(n.Left)
		b.DFSPreOHelper(n.Right)
	}
}
//explore left of current node, explore right of current node then look at node
func (b *BST) DFSPostO(){
	b.DFSPostOHelper(b.Root)
	fmt.Println("")
}

func (b *BST) DFSPostOHelper(n *Node){
	if n == nil {
		return
	} else {
		b.DFSPostOHelper(n.Left)
		b.DFSPostOHelper(n.Right)
		fmt.Print(n.Value, " ")
	}
}

//explore left of current node, explore current node then explore right of current node -- get tree in ascending order
func (b *BST) DFSInO(){
	b.DFSInOHelper(b.Root)
	fmt.Println("")
}

func (b *BST)  DFSInOHelper(n *Node){
	if n == nil {
		return
	} else {
		b.DFSInOHelper(n.Left)
		fmt.Print(n.Value, " ")
		b.DFSInOHelper(n.Right)
	}
}