package main

import "fmt"

func main(){
	fmt.Println(divide(100000000000,-2))
}
/*
	* Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

	* Return the quotient after dividing dividend by divisor.

	* The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.
*/
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	resultIsNegative := false 
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0){
		resultIsNegative = true
	}
	dividend = cTP(dividend)
	divisor = cTP(divisor)

	running := divisor
	result := 0

	for running <= dividend {
		result++
		running += divisor
	}

	if resultIsNegative {
		return cTN(result)
	}
	return result
}

func cTP(n int) int {
	if n >= 0 {
		return n
	}
	i := n 
	for i + n != 0 {
		i++
	}
	return i
}

func cTN(n int) int {
	if n <= 0 {
		return n
	}
	i := n 
	for i + n != 0 {
		i--
	}
	return i
}