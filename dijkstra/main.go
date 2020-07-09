package main 

import (
	"fmt"
	"errors"
)

func main(){
	g := &Graph{}
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")
	g.AddVertex("F")
	g.AddUndirectedEdge("A","B", 4)
	g.AddUndirectedEdge("A","C", 2)
	g.AddUndirectedEdge("B","E", 3)
	g.AddUndirectedEdge("C","D", 2)
	g.AddUndirectedEdge("C","F", 4)
	g.AddUndirectedEdge("D","F", 1)
	g.AddUndirectedEdge("D","E", 3)
	g.AddUndirectedEdge("F","E", 1)
	g.Print()
	fmt.Println(g.ShortestPath("A","E"))
}

type PriorityQueue struct {
	Size int
	Values []*QueueNode
}

type QueueNode struct {
	Vertex string 
	Weight int
}

func (h *PriorityQueue) Enqueue(vertex string, weight int) error {
	if h.Size == len(h.Values) {
		return errors.New("Heap is full, cannot add new value")
	}
	v := &QueueNode{Vertex: vertex, Weight: weight}
		//put value at the end of array
	currentIndex := h.Size
	h.Size++
	h.Values[currentIndex] = v 
	if currentIndex == 0 {
		return nil
	}
	//check if it has parent
	parentIndex := (currentIndex - 1) / 2
		//if it has, compare with parent
	for currentIndex > 0 && h.Values[currentIndex].Weight < h.Values[parentIndex].Weight{
		//if greater than parent, swap
		tmp := h.Values[parentIndex]
		h.Values[parentIndex] = h.Values[currentIndex]
		h.Values[currentIndex] = tmp
		currentIndex = parentIndex 
		parentIndex = (currentIndex - 1) / 2
		//repeat until it is either less than parent or has no parent
	}
	return nil
}

func (h *PriorityQueue) Dequeue() (*QueueNode,error) {
	if h.Size == 0 {
		return nil,errors.New("Cannot remove from an empty heap")
	}
	//get root value
	v := h.Values[0] 
	//remove root and restructure tree

	//replace root with last value in array
	h.Values[0] = h.Values[h.Size - 1]
	//initialize index of last value to 0
	h.Values[h.Size - 1] = nil
	h.Size--
	currentIndex := 0
	//sift the last value down -- I use less than because if it is equal to size we are already at the last
	for currentIndex < h.Size {
		rI := (2 * currentIndex) + 2
		lI := (2 * currentIndex) + 1
		compareIndex := rI 
		//choose the child that is smallest and exists
		if rI > h.Size - 1 && lI > h.Size - 1 {
			break
		} else if rI > h.Size - 1 {
			compareIndex = lI 
		} else if lI > h.Size - 1 {
			compareIndex = rI 
		} else {
			if h.Values[lI].Weight < h.Values[rI].Weight {
				compareIndex = lI
			}
		}
		//If it is greater than child swap (we want a min heap)
		if h.Values[currentIndex].Weight > h.Values[compareIndex].Weight {
			tmp := h.Values[compareIndex]
			h.Values[compareIndex] = h.Values[currentIndex]
			h.Values[currentIndex] = tmp
			currentIndex = compareIndex
		} else {
			break
		}
	}
	//return root value
	return v, nil
}

/**
	* Vertex - node 
	* Edge - connection between nodes
	*/
type Graph struct {
	AdacencyList map[string]map[string]int
} 

func (g *Graph) Print() {
	for k,v := range g.AdacencyList {
		fmt.Print(k,": ")
		for e,w := range v {
			fmt.Print(e,"->",w," ")
		}
		fmt.Println("")
	}
}

func (g *Graph) AddVertex(name string) {
	if g.AdacencyList == nil {
		g.AdacencyList = make(map[string]map[string]int)
	}
	if _,ok := g.AdacencyList[name]; !ok {
		//only add vertex if it does not exist
		g.AdacencyList[name] = make(map[string]int)
	}
}

func (g *Graph) AddUndirectedEdge(source, destination string, weight int) error {
	//check if source and destination exist
	if _,ok := g.AdacencyList[source]; !ok {
		return errors.New("Vertex " + source + " does not exist")
	}

	if _,ok := g.AdacencyList[destination]; !ok {
		return errors.New("Vertex " + destination + " does not exist")
	}

	g.AdacencyList[source][destination] = weight
	g.AdacencyList[destination][source] = weight
	return nil
}

func (g *Graph) InitializeQueue(capacity int) *PriorityQueue {
	m := &PriorityQueue{
		Values: make([]*QueueNode,capacity),
	}
	return m
}

/**
	* Dijkstra's shortest path algorithm
	*/
func (g *Graph) ShortestPath(source, destination string) []string{
	const MaxInt = int(^uint(0)>>1)
	result := make([]string,0)
	//visited := make(map[string]bool)//all vertices I have visited previosuly
	queue := g.InitializeQueue(255)//prioritize vertex with smallest path
	previous := make(map[string]string)//the vertice I previously visited from current vertice
	distances := make(map[string]int)//ditance from source to vertice [string]

	//loop through all vertices of the graph
	for vertex,_ := range g.AdacencyList {
		if source == vertex {
			distances[vertex] = 0
			queue.Enqueue(vertex, 0)
		} else {
			distances[vertex] = MaxInt
			queue.Enqueue(vertex, MaxInt)
		}
		previous[vertex] = ""
	}

	for queue.Size > 0 {
		smallest, err := queue.Dequeue()
		if err == nil {
			if smallest.Vertex == destination {
				fmt.Println(previous)
				//destination reached
				currentVertex := smallest.Vertex 
				fmt.Println(currentVertex, previous[currentVertex])
				for previous[currentVertex] != "" {
					result = append(result, currentVertex)
					currentVertex = previous[currentVertex]
				}
				break
			}
			if smallest != nil || distances[smallest.Vertex] != MaxInt {
				//get neighbors
				for neighbor,weight := range g.AdacencyList[smallest.Vertex] {
					//find smallest distance to neighbor
					candidate := distances[smallest.Vertex] + weight
					if candidate < distances[neighbor] {
						distances[neighbor] = candidate
						previous[neighbor] = smallest.Vertex
						queue.Enqueue(neighbor, candidate)
					}
				}
			}
		} else {
			fmt.Println(err)
		}
	}
	result = append(result,source)
	//reverse result 
	n := len(result)
	for i := n - 1; i  >= n / 2; i-- {
		in := result[n - i - 1]
		result[n - i - 1] = result[i]
		result[i] = in
	}
	return result
}