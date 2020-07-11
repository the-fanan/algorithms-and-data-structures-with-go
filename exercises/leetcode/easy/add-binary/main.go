package main

import (
	"fmt"
)

func main(){
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a,b string) string {
	result := ""
	n := len(a)
	m := len(b)
	carry := 0 
	for n > 0 || m > 0 || carry > 0 {
		if n > 0 && m > 0 {
			s := d(a[n - 1]) + d(b[m - 1]) + carry
			if s > 1 {
				carry = 1
				res := s % 2
				result = prepend(res, result)
			} else {
				carry = 0
				result = prepend(s, result)
			}
		} else if n > 0 {
			s := d(a[n - 1])  + carry
			if s > 1 {
				carry = 1
				res := s % 2
				result = prepend(res, result)
			} else {
				carry = 0
				result = prepend(s, result)
			}
		} else if m > 0 {
			s := d(b[m - 1]) + carry
			if s > 1 {
				carry = 1
				res := s % 2
				result = prepend(res, result)
			} else {
				carry = 0
				result = prepend(s, result)
			}
		} else {
			//carry has to be the last man standing
			result = prepend(carry, result)
			carry = 0
		}
		n--
		m--
	}
	return result
}

func d(n byte) int{
	return int(n) - 48
}

func prepend(n int, r string) string {
	inter := ""
	inter += string(byte(n + 48))
	inter += r 
	r = inter
	return r
}