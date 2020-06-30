package main

import "fmt"

func main() {
	fmt.Println(SelectionSort([]int{5,7,1,3,9,3,7,9,8,3,4,6,1,2,7,10,5,7,11,6,1}))
}

func SelectionSort(a []int) []int{
	n := len(a)
	//at each index, loop through array and look for value that is the least the swap with item at current index.
	for i := 0; i < n; i++{
		minIdx := i//initalize minimum value
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		//we have the min index now swap it with a[i]
		if minIdx != i {
			in := a[i]
			a[i] = a[minIdx]
			a[minIdx] = in
		}
	}
	return a
}