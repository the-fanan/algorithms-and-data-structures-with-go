package main

import (
	"fmt"
)
//time related functions
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

//space related functions
func ReverseString(s string) string {
	//space of O(n)
	result := ""
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		result = result + string(s[i])
	}
	return result
}
func Sum() int{
	//space O(1)
	a := []int{1,2,3,4,5,6}
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func ReverseStringC(s []byte) []byte {
	//space of O(1)
	n := len(s)
	for i := n - 1; i >= n / 2; i-- {
		in := s[n - i - 1]
		s[n - i - 1] = s[i]
		s[i] = in
	}
	return s
}


func main(){
	PrintAllPairs(1)
	fmt.Println(ReverseStringC([]byte("Fanan")))
}
/**
	* For simplifying Big O
	*
	* 1. Constants do not matter. They should always be equivalent to 1
	* 2. Smaller terms do not matter. If you have O(n^2) and O(n) (n^2 + n) the notation should be O(n^2)
*/

/**
	* For Space
	* Primitive (int, bool, nil) types take up constant space O(1)
	* Strings, arrays, slice, maps, structs require O(n) space
*/
