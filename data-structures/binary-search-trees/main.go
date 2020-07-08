package main

import (
	"fmt"
	_"errors"
)

func main(){
	b := &BST{}
	b.Insert(10)//root
	b.Insert(13)
	b.Insert(5)
	b.Insert(16)
	b.Insert(12)
	b.Insert(6)
	b.Insert(4)
	fmt.Println(*b.Root)
	fmt.Println(b.Find(6))
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