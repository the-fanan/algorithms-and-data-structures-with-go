package main

import (
	"fmt"
)

func main(){
	a := []int{4,6,2,7,8,90,0,64,1,42,13,16}
	sa := []int{1,2,3,4,5,8,9,10}
	fmt.Println(linearSearch(a,64))
	fmt.Println(binarySearch(sa,9))
}

func linearSearch(a []int, v int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == v {
			return i
		}
	}
	return -1
}

func binarySearch(a []int, v int) int {
	l := 0
	r := len(a) - 1
	m := (l + r) / 2
	for l <= r {
		if a[m] == v {
			return m
		}else {
			if a[m] > v {
				r = m - 1
			} else {
				l = m + 1
			}
			m = (l + r) / 2
		}
	}
	return -1
}