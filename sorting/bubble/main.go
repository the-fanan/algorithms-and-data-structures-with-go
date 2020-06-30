package main

import "fmt"

func main(){
	fmt.Println(BubbleSort([]int{5,7,1,3,9,3,7,9,8,3,4,1,2,7,10,5,7,11,230}))
}

func BubbleSort(a []int) []int{
	n := len(a)
	swapped := true
	for swapped {
		swaps := 0
		for i := 0; i < n - 1; i++ {
			if a[i] > a[i + 1]{
				in := a[i]
				a[i] = a[i + 1]
				a[i + 1] = in
				swaps++
			}
		}
		if swaps == 0 {
			swapped = false
		}
	}
	return a
}