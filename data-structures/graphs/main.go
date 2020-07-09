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
	fmt.Println("--------------Remove Undirected Edge--------------")
	g.RemoveUndirectedEdge("lagos", "oyo")
	g.Print()
	fmt.Println("--------------Remove Vertex--------------")
	g.RemoveVertex("lagos")
	g.Print()
	fmt.Println("--------------Add Vertex and Edges for traversal--------------")
	g.AddVertex("lagos")
	g.AddUndirectedEdge("lagos", "abuja")
	g.AddUndirectedEdge("lagos", "oyo")
	g.AddUndirectedEdge("oyo", "plateau")
	g.AddUndirectedEdge("kano", "abuja")
	g.Print()
	fmt.Println("--------------Depth First Traversal--------------")
	g2 := &Graph{}
	g2.AddVertex("A")
	g2.AddVertex("B")
	g2.AddVertex("C")
	g2.AddVertex("D")
	g2.AddVertex("E")
	g2.AddVertex("F")
	g2.AddUndirectedEdge("A", "C")
	g2.AddUndirectedEdge("A", "B")
	g2.AddUndirectedEdge("B", "D")
	g2.AddUndirectedEdge("D", "E")
	g2.AddUndirectedEdge("D", "F")
	g2.AddUndirectedEdge("E", "F")
	g2.AddUndirectedEdge("C", "E")
	fmt.Println(g2.DFTraversal("A"))
	fmt.Println("--------------Breadth First Traversal--------------")
	g3 := &Graph{}
	g3.AddVertex("A")
	g3.AddVertex("B")
	g3.AddVertex("C")
	g3.AddVertex("D")
	g3.AddVertex("E")
	g3.AddVertex("F")
	g3.AddUndirectedEdge("A", "E")
	g3.AddUndirectedEdge("A", "B")
	g3.AddUndirectedEdge("B", "D")
	g3.AddUndirectedEdge("B", "C")
	g3.AddUndirectedEdge("C", "D")
	g3.AddUndirectedEdge("D", "F")
	g3.AddUndirectedEdge("D", "E")
	g3.AddUndirectedEdge("F", "E")
	fmt.Println(g3.BFTraversal("A"))
	fmt.Println("g2 BFT:", g2.BFTraversal("A"))
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

func (l *List) Get(p int) (*ListNode,error) {
	if p < 0 || p > l.Length - 1 {
		err := errors.New("Index is out of range")
		return nil,err
	}
	i := 0
	n := l.Head
	for i < p {
		n = n.Next
		i++
	}
	return n,nil
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
				l.Shift()
			} else {
				if current.Next == nil {
					l.Tail = prev
				} else {
					prev.Next = current.Next
				}
				current.Next = nil
				l.Length--
			}
			break
		}
	}
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

/**
	* Traversing a graph
	*
	* Depth First
	* Breadth First
	*/

func (g *Graph) DFTraversal(origin string) []string {
	stack := make([]string,0)//push append(stack, value) // pop stack[:len(stack) - 1]
	result := make([]string,0)
	visited := make(map[string]bool)
	stack = append(stack, origin)
	for len(stack) > 0{
		//pop vertex from stack
		vertex := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		//if vertex has not been visited
		if _,ok := visited[vertex]; !ok {
			//add to visited vertexes
			visited[vertex] = true 
			//add to our results array to show route
			result = append(result, vertex)
			//loop through edges for vertex and add them to stack
			for i := 0; i < g.AdacencyList[vertex].Length; i++{
				vert,_ := g.AdacencyList[vertex].Get(i)
				stack = append(stack, vert.Value)
			}
		}
	}
	return result
}

func (g *Graph) BFTraversal(origin string) []string {
	queue := make([]string,0)//enqueue append(stack, value) // dequeue stack[1:]
	result := make([]string,0)
	visited := make(map[string]bool)
	queue = append(queue, origin)
	for len(queue) > 0{
		//pop vertex from stack
		vertex := queue[0]
		queue = queue[1:]
		//if vertex has not been visited
		if _,ok := visited[vertex]; !ok {
			//add to visited vertexes
			visited[vertex] = true 
			//add to our results array to show route
			result = append(result, vertex)
			//loop through edges for vertex and add them to queue
			for i := 0; i < g.AdacencyList[vertex].Length; i++{
				vert,_ := g.AdacencyList[vertex].Get(i)
				queue = append(queue, vert.Value)
			}
		}
	}

	return result
}