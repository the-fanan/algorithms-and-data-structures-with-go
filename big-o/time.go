package main

import (
	"fmt"
)

func reverseString(s string) string {
	result := ""
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		result = result + string(s[i])
	}
	return result
}

/**
	* Which between the sum functions is more efficient?
	*
	* SumNumbers1 has running time of O(n)
	* SumNumbers2 has running time of O(1)
	*
	* SumNumbers2 is more efficient
*/
func SumNumbers1(n int) int{
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	return s
}

func SumNumbers2(n int) int{
	return n * (n + 1)/2
}

/**
	* This function has a run time of O(n^2)
	* O(n) inside O(n)
*/
func PrintAllPairs(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			fmt.Println(i,j)
		}
	}
}


func main(){
	PrintAllPairs(1)
}
/**
	* For simplifying Big O
	*
	* 1. Constants do not matter. They should always be equivalent to 1
	* 2. Smaller terms do not matter. If you have O(n^2) and O(n) (n^2 + n) the notation should be O(n^2)
*/
