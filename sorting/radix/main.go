package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(RadixSort([]int{5,7,1,3,9,3,7,9,8,3,4,6,1,2,7,10,5,321,7,11,6,1,234}))
}
//for positive numbers
func RadixSort(a []int) []int{
	n := len(a)
	m := 10
	//find the largest number
	max := 0
	for _,v := range a {
		if v > max {
			max = v
		}
	}
	//get number of times loop needs to run
	lr := 1
	if max > 0 {
		lr = int(math.Floor(math.Log10(float64(max)))) + 1//how many digits does this number have?
	}

	for k := 0; k < lr; k++{
		//mote: update bucket to use queues next time
		bucket := make([][]int,10)
		for i := 0; i < n; i++{
			ind := int(math.Floor(float64(a[i]) / math.Pow(10, float64(k)))) % 10//get the kth digit from behind to use as index of bucket
			bucket[ind] = append(bucket[ind], a[i])
		}

		ai := 0
		for _,b := range bucket{
			for _,v := range b{
				a[ai] = v
				ai++
			}
		}
		m *= 10
	}
	return a
}