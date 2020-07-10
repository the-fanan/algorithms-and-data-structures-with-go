package main

import (
	"fmt"
)

func main(){
	fmt.Println(plusOne([]int{9,9,9}))
}

/**
	* Given a non-empty array of digits representing a non-negative integer, increment one to the integer.
	*
	* The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.
	*
	* You may assume the integer does not contain any leading zero, except the number 0 itself.
*/
func plusOne(digits []int) []int {
	carry := 0
	n := len(digits)
	for i := n - 1; i >= 0; i--{
		s := 0
		if i == n - 1 {
			s = digits[i] + carry + 1
		} else {
			s = digits[i] + carry
		}
		
		if s >= 10 {
			carry = s / 10
			res := s % 10
			digits[i] = res
		} else {
			carry = 0
			digits[i] = s
		}

		if carry == 0 {
			break
		}
	}
	
	if carry > 0 {
		ln := make([]int, n + 1)
		ln[0] = carry 
		for i,v := range digits {
			ln[i + 1] = v
		}
		return ln
	}
	
	return digits
}