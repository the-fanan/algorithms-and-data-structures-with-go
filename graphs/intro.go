package main

import (
	"fmt"
)

type Vertex struct {
	ID string
}

func (v *Vertex) Element() string {
	return ""
}

type Edge struct {
	Name string //name of the street
	Distance float64 //how far is it
	Source Vertex 
	Destination Vertex
}

func (e *Edge) Enpoints() (Vertex, Vertex) {
	return e.Source, e.Destination
}

func (e *Edge) Opposite(v Vertex) Vertex {
	if v.ID == e.Source.ID {
		return e.Destination
	} 
	return e.Source
}

type Graph struct {

}

func (g *Graph) VertexCount() int {
	return 0
}

func (g *Graph) EdgeCount() int {
	return 0
}

func (g *Graph) GetEdge(source, destination *Vertex) *Edge {
	return nil
}

func (g *Graph) InsertVertex() *Vertex {
	return nil
}

func (g *Graph) InsertEdge(source, destination *Vertex) *Edge {
	return nil
}

func (g *Graph) RemoveVertex(v *Vertex) bool {
	return true
}

func (g *Graph) RemoveEdge(e *Edge) bool {
	return true
}

//adjacency map for representing graph


func main(){
	fmt.Println("Hello World");
}