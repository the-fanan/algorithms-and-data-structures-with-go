package main

import 
(
	"fmt"
	"errors"
)

func main(){
	q := &Queue{} 
	q.Enqueue(5)
	q.Enqueue(10)
	q.Enqueue(15)
	q.Print()
	q.Dequeue()//5
	q.Print()
	q.Dequeue()//10
	q.Print()
	q.Dequeue()//15
	q.Print()
}

type Node struct {
	Next *Node 
	Value int
}

/**
	* Follows First In First Out principle (FIFO)
	*/
type Queue struct {
	Head *Node 
	Tail *Node
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

func (q *Queue) Dequeue() (int,error) {
	if q.Length == 0 {
		return 0, errors.New("Queue is empty")
	}
	node := q.Head 
	q.Head = q.Head.Next 
  node.Next = nil
	q.Length--
	return node.Value, nil
}

func (q *Queue) Enqueue(v int) {
	n := &Node{Value: v}
	if q.Length == 0 {
		q.Head = n
	} else {
		q.Tail.Next = n
	}
	q.Tail = n
	q.Length++
}