package main

import (
	"fmt"
	"errors"
)

func main(){
	g := &Graph{}
	g.AddVertex("lagos")
	g.AddVertex("abuja")
	g.AddVertex("oyo")
	g.AddVertex("kano")
	g.AddVertex("plateau")
	g.AddUndirectedEdge("lagos", "abuja")
	g.AddUndirectedEdge("lagos", "oyo")
	g.AddUndirectedEdge("plateau", "kano")
	g.AddUndirectedEdge("lagos", "plateau")
	g.Print()
	fmt.Println("--------------Remove Undirected Edge-----------------")
	g.RemoveUndirectedEdge("lagos", "oyo")
	g.Print()
	fmt.Println("--------------Remove Vertex-----------------")
	g.RemoveVertex("lagos")
	g.Print()
}
type ListNode struct {
	Value string 
	Next *ListNode
}

type List struct {
	Head *ListNode 
	Tail *ListNode 
	Length int 
}

func (l *List) Print(){
	if l.Head == nil{
		fmt.Print("List is empty")
	} else {
		current := l.Head
		i := 0
		for current != nil {
			fmt.Print(i, ": ", current.Value, ", ")
			current = current.Next
			i++
		}
	}
	fmt.Println("")
}

func (l *List) Push(v string) {
	n := &ListNode{Value: v}
	if (l.Head == nil) {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
	l.Length++
}

func (l *List) Shift() string {
	if l.Head == nil {
		return ""
	}
	current := l.Head 
	if l.Head == l.Tail {
		l.Tail = nil
		l.Head = nil
	} else {
		l.Head = l.Head.Next
	}
	current.Next = nil
	l.Length--
	return current.Value
}

func (l *List) Remove(v string){
	if l.Head == nil {
		return
	}
	prev := &ListNode{}
	prev = nil 
	current := l.Head 
	for current != nil {
		if current.Value != v {
			prev = current 
			current = current.Next
		} else {
			if prev == nil {
				//we are at the head 
				if l.Head == l.Tail {
					l.Head = nil
					l.Tail = nil
				} else {
					l.Head = current.Next
				}
			} else {
				prev.Next = current.Next
			}
			current.Next = nil
			break
		}
	}
	l.Length--
}
/**
	* Vertex - node 
	* Edge - connection between nodes
	*/
type Graph struct {
	AdacencyList map[string]*List
} 

func (g *Graph) Print() {
	for k,v := range g.AdacencyList {
		fmt.Print(k,": ")
		v.Print()
	}
}

func (g *Graph) AddVertex(name string) {
	if g.AdacencyList == nil {
		g.AdacencyList = make(map[string]*List)
	}
	if _,ok := g.AdacencyList[name]; !ok {
		//only add vertex if it does not exist
		g.AdacencyList[name] = &List{}
	}
}

func (g *Graph) AddUndirectedEdge(source, destination string) error {
	//check if source and destination exist
	if _,ok := g.AdacencyList[source]; !ok {
		return errors.New("Vertex " + source + " does not exist")
	}

	if _,ok := g.AdacencyList[destination]; !ok {
		return errors.New("Vertex " + destination + " does not exist")
	}

	g.AdacencyList[source].Push(destination)
	g.AdacencyList[destination].Push(source)
	return nil
}

func (g *Graph) AddDirectedEdge(source, destination string) error {
	//check if source and destination exist
	if _,ok := g.AdacencyList[source]; !ok {
		return errors.New("Vertex " + source + " does not exist")
	}

	if _,ok := g.AdacencyList[destination]; !ok {
		return errors.New("Vertex " + destination + " does not exist")
	}
	g.AdacencyList[source].Push(destination)
	return nil
}

func (g *Graph) RemoveUndirectedEdge(source, destination string) error {
	if _,ok := g.AdacencyList[source]; !ok {
		return errors.New("Vertex " + source + " does not exist")
	}

	if _,ok := g.AdacencyList[destination]; !ok {
		return errors.New("Vertex " + destination + " does not exist")
	}
	g.AdacencyList[source].Remove(destination)
	g.AdacencyList[destination].Remove(source)
	return nil
}

func (g *Graph) RemoveDirectedEdge(source, destination string) error {
	if _,ok := g.AdacencyList[source]; !ok {
		return errors.New("Vertex " + source + " does not exist")
	}

	if _,ok := g.AdacencyList[destination]; !ok {
		return errors.New("Vertex " + destination + " does not exist")
	}
	g.AdacencyList[source].Remove(destination)
	return nil
}

func (g *Graph) RemoveVertex(vertex string) error {
	if _,ok := g.AdacencyList[vertex]; !ok {
		return errors.New("Vertex " + vertex + " does not exist")
	}
	for g.AdacencyList[vertex].Length > 0 {
		dest := g.AdacencyList[vertex].Shift()
		g.RemoveDirectedEdge(dest, vertex)
	}
	delete(g.AdacencyList, vertex)
	return nil
}