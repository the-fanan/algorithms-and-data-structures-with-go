package main

import (
	"fmt"
	"errors"
)

func main(){
	m := InitializeHeap(15)
	m.Insert(100)
	fmt.Println(m.Values)
	m.Insert(50)
	fmt.Println(m.Values)
	m.Insert(60)
	fmt.Println(m.Values)
	m.Insert(70)
	fmt.Println(m.Values)
	m.Insert(80)
	fmt.Println(m.Values)
	m.Insert(30)
	fmt.Println(m.Values)
	m.Insert(40)
	fmt.Println(m.Values)
	m.Insert(90)
	fmt.Println(m.Values)
	m.Insert(35)
	fmt.Println(m.Values)
	m.Insert(55)
	fmt.Println(m.Values)
	m.Insert(68)
	fmt.Println(m.Values)
	m.Insert(45)
	fmt.Println(m.Values)
	m.Insert(76)
	fmt.Println(m.Values)
	m.Insert(20)
	fmt.Println(m.Values)
	m.Insert(10)
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
	m.Extract()
	fmt.Println(m.Values)
}
/**
	* For a max Heap, parent must be greater than both children.
	*
	* Used for Priority queues and to aid in Graph Traversals
	* li = 2n + 1 --> index of Left child --- n = floor[(li - 1) / 2]
	* ri = 2n + 2 --> index of Right child --- n = floor[(ri - 1) / 2]
	* Richt indexed are always even, left indexes are always odd
	*/
type MaxHeap struct {
	Size int
	Values []int
}

func InitializeHeap(capacity int) *MaxHeap {
	m := &MaxHeap{
		Values: make([]int,capacity),
	}
	return m
}

func (h *MaxHeap) Insert(v int) error{
	if h.Size == len(h.Values) {
		return errors.New("Heap is full, cannot add new value")
	}
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
	for currentIndex > 0 && h.Values[currentIndex] > h.Values[parentIndex]{
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

func (h *MaxHeap) Extract() (int,error) {
	if h.Size == 0 {
		return 0,errors.New("Cannot remove from an empty heap")
	}
	//get root value
	v := h.Values[0] 
	//remove root and restructure tree

	//replace root with last value in array
	h.Values[0] = h.Values[h.Size - 1]
	//initialize index of last value to 0
	h.Values[h.Size - 1] = 0
	h.Size--
	currentIndex := 0
	//sift the last value down -- I use less down because if it is equal to size we are already at the last
	for currentIndex < h.Size {
		rI := (2 * currentIndex) + 2
		lI := (2 * currentIndex) + 1
		compareIndex := rI 
		if rI > len(h.Values) - 1 && lI > len(h.Values) - 1 {
			break
		} else if rI > len(h.Values) - 1 {
			compareIndex = lI 
		} else if lI > len(h.Values) - 1 {
			compareIndex = rI 
		} else {
			if h.Values[lI] > h.Values[rI] {
				compareIndex = lI
			}
		}

		if h.Values[currentIndex] < h.Values[compareIndex] {
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