package main

import "fmt"

func main(){
	fmt.Println(MergeSort([]int{5,7,1,3,9,3,7,9,8,3,4,6,1,2,7,10,5,7,11,6,1}))
}

func MergeSort(a []int) []int{
	n := len(a)
	ta := make([]int, n)
	recMergeSort(&a, &ta,0,n - 1)
	return a
}

//left index, mid+ 1, end r+1
func Merger(a *[]int,ta *[]int,l,r,end int){
	left := l
	right := r 
	t := l
	
	for left < r && right < end {
		if (*a)[left] < (*a)[right] {
			(*ta)[t] = (*a)[left]
			left++
		} else {
			(*ta)[t] = (*a)[right]
			right++
		}
		t++
	}

	for left < r {
		(*ta)[t] = (*a)[left]
		t++
		left++
	}

	for right < end {
		(*ta)[t] = (*a)[right]
		t++
		right++
	}

	for i := l; i < end; i++{
		(*a)[i] = (*ta)[i]
	}
}

func recMergeSort(a *[]int, ta *[]int,l,r int){
	if l == r {
		return
	} else {
		m := (l + r) / 2
		recMergeSort(a, ta, l, m)
		recMergeSort(a, ta, m+1, r)
		Merger(a, ta, l, m+1, r+1)
	}
}