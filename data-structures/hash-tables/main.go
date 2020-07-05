package main

import (
	"fmt"
	"errors"
)

/**
	* Qualities of a good hash
	* 1. Fast O(1) time
	* 2. Deterministic
	* 3. Uniform distribution
	*/
func main(){
	h := &HashTable{
		Size: 100,
		Container: make([]*List, 100),
	}
	h.Set("black","#000")
	v, e := h.Get("black")
	fmt.Println(v, e)
	v, e = h.Get("white")
	fmt.Println(v, e)
	h.Set("black","#0err")
	v, e = h.Get("black")
	fmt.Println(v, e)
	h.Set("white","#fff")
	v, e = h.Get("white")
	fmt.Println(v, e)
	h.Set("pink","#fff")
	v, e = h.Get("white")
	fmt.Println(h.Values())
	fmt.Println(h.Keys())
}

//this hash uses separate chainging
type HashTable struct {
	Container []*List
	Size int
}

func (h *HashTable) Values() []string {
	values := make([]string,0)
	for _, list := range h.Container {
		if list != nil {
			for i:= 0; i < list.Length; i++ {
				node, err := list.GetByIndex(i) 
				if err != nil {
					fmt.Println(err)
					continue
				}
				values = append(values,node.Value)
			}
		}
	}
	//ensure values are unique 
	indexHolder := make(map[string]int)
	for i, value := range values {
		indexHolder[value] = i
	}
	result := make([]string,0)
	for key,_ := range indexHolder {
		result = append(result, key)
	}
	return result
}

func (h *HashTable) Keys() []string {
	keys := make([]string,0)
	for _, list := range h.Container {
		if list != nil {
			for i:= 0; i < list.Length; i++ {
				node, err := list.GetByIndex(i) 
				if err != nil {
					fmt.Println(err)
					continue
				}
				keys = append(keys,node.Key)
			}
		}
	}
	return keys
}

func (h *HashTable) Get(key string) (string, error) {
	hash := h.Hash(key)
	if h.Container[hash] == nil {
		return "", errors.New("Key does not exist")
	}
	l := h.Container[hash]
	node, err := l.Get(key)
	if err != nil {
		return "", errors.New("Key does not exist")
	}
	return node.Value, nil
}

func (h *HashTable) Set(key, value string) {
	hash := h.Hash(key)
	l := &List{} 
	if h.Container[hash] == nil {
		l.Push(key,value)
		h.Container[hash] = l
	} else {
		l = h.Container[hash]
		//ensure keys are unique 
		node, err := l.Get(key)
		if err == nil {
			//key exists so update it
			node.Value = value
		} else {
			//key does not exist so create new one
			l.Push(key,value)
		}
	}
}

func (h *HashTable) Hash(key string) int{
	n := 100
	if len(key) < n {
		n = len(key)
	}
	const PRIME = 97
	sum := 0 
	for i := 0; i < n; i++{
		sum = (sum * PRIME + int(key[i]) - 96) % h.Size
	}
	return sum
}
/**
	* Handling Collisions
	* 1. Separate Chaining - store in same index with help of another data structure
	* 2. Linear Probing - bounce off index to neearest empty index
	*/
//String hash s0 a^n−1 + s1 a^n−2 + · · · + sn−3 a^2 + sn−2 a + sn−1
// a is a Constant, s is the ASCII integer equivalent of the string
/**
	* PROBES
  * i is the iTH probe in the sequence, home is the position the key was originaly mapped to
 	* Liniear = (home + i) % N
	* Modified Linear = (home + i ∗ c) % N
	* Quadratic = (home + i^2 ) % M
	* Double hashing = (home + i ∗ hp(key)) % M where hp(key) = 1 + key % P where P < M and M & P are prime
*/

//list to implement separate chainging
type Node struct{
	Value string 
	Key string //original key
	Next *Node
}

type List struct {
	Length int 
	Head *Node 
	Tail *Node
}

func (l *List) Print() {
	current := l.Head 
	for current != nil {
		fmt.Println("key:", current.Key, "value:", current.Value)
		current = current.Next
	}
}

func (l *List) Push(key string, value string) {
	n := &Node{Key: key, Value: value}
	if l.Length == 0 {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n 
	l.Length++
}

func (l *List) Get(key string) (*Node, error) {
	current := l.Head 
	for current != nil {
		if current.Key == key {
			break;
		}
		current = current.Next
	}
	if current == nil {
		return nil, errors.New("Key does not exist")
	}
	return current,nil
}

func (l *List) GetByIndex(index int) (*Node, error) {
	if index < 0 || index > l.Length - 1 {
		err := errors.New("Index is out of range")
		return nil,err
	}
	i := 0
	n := l.Head
	for i < index {
		n = n.Next
		i++
	}
	return n,nil
}

