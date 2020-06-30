package main

import "fmt"

func main(){
	fmt.Println(QuickSort([]int{5,7,1,3,9,3,7,9,8,3,4,6,1,2,7,10,5,7,11,6,1}))
}

func QuickSort(a []int) []int{
	n := len(a)
	recQS(&a, 0, n - 1)
	return a
}

func recQS(a *[]int, start, end int) {
	if start >= end {
		return
	} else {
		p := pivot(a, start, end + 1)
		recQS(a,start, p - 1)
		recQS(a,p + 1, end)
	}
}

func pivot(a *[]int, start, end int) int{
	pivot := (*a)[start]
	swapIndex := start
	for i := start + 1; i < end; i++{
		if pivot > (*a)[i] {
			swapIndex++
			in := (*a)[swapIndex]
			(*a)[swapIndex] = (*a)[i]
			(*a)[i] = in
		}
	}
	in := (*a)[swapIndex]
	(*a)[swapIndex] = (*a)[start]
	(*a)[start] = in
	return swapIndex
}