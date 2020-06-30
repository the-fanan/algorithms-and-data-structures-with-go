package main

import "fmt"

func main(){
	fmt.Println(InsertionSort([]int{5,7,1,3,9,3,7,9,8,3,4,6,1,2,7,10,5,7,11,6,1}))
}
//best for arays that are almost sorted
//can be used for continuous addtion sort
func InsertionSort(a []int) []int{
	n := len(a)
	//at each index loop backwards till you find where it is the least
	for i := 1; i < n; i++{
		j := i
		for j > 0 && a[j - 1] > a[j] {
			in := a[j - 1]
			a[j - 1] = a[j]
			a[j] = in
			j--
		}
	}
	return a
}