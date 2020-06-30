package main

import (
	"fmt"
)

func main(){
	fmt.Println(CountSort([]int{5,7,1,3,9,3,7,9,8,3,4,6,1,2,7,10,5,7,11,6,1}))
}

//can oly be used for positive integers
//can be modified for negative but would add complexity to code
func CountSort(a []int) []int{
	n := len(a)
	max := 0
	for _,v := range a {
		if v > max {
			max = v
		}
	}
	ca := make([]int, max + 1)
	for i := 0; i < n; i++{
		ca[a[i]]++
	}
	ai := 0
	for i := 0; i < max + 1; i++{
		for ca[i] > 0{
			a[ai] = i 
			ca[i]--
			ai++
		}
	}

	return a
}