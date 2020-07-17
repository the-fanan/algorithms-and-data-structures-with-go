package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{89,62,70,58,47,47,46,76,100,70}))
}

/*
	* Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.

	* For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

	* Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100]. 
*/	

func dailyTemperatures(T []int) []int {
	temps := make([][]int,101)
	for i,v := range temps {
		if v == nil {
			temps[i] = make([]int,0)
		} 
	}

	for i := len(T) - 1; i >= 0; i-- {
		temps[T[i]] = append(temps[T[i]],i)
	}

	for i,v := range T {
		fI := -1 
		for j := v + 1; j < len(temps); j++ {
			for _,idx := range temps[j] {
				if idx > i {
					if fI == -1 {
						fI = idx
					}	else if fI > idx {
						fI = idx
					}
				}
			}
		}

		if fI == -1 {
			T[i] = 0
		} else {
			T[i] = fI - i
		}
	}

	return T
}