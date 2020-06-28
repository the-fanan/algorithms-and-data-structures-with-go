package main

import (
	"fmt"
)

func main() {
	a := []int{2,4,1,8}
	fmt.Println(someRecursive(a, isOdd))
}

func reverseString(s string) string{
	b := []byte(s)
	rSH(&b, len(s) - 1)
	return string(b)
}

func rSH(s *[]byte, i int){
	n := len(*s)
	if i < (n/2) {
		return
	} else {
		in := (*s)[n - i - 1]
		(*s)[n - i - 1] = (*s)[i]
		(*s)[i] = in
		rSH(s, i-1)
	}
}

func isPalindrome(s string) bool {
	return iPH(s,len(s) - 1)
}

func iPH(s string, i int) bool{
	n := len(s)
	if i == n / 2 {
		//this is the base point
		if n % 2 == 0 {
			if s[n - 1 - i] == s[i] {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		if s[n - 1 - i] == s[i] {
			return iPH(s, i - 1)
		} else {
			return false
		}
	}
}

type callback func (i int) bool 

func isOdd(i int) bool {
	if i % 2 == 0 {
		return false
	} else {
		return true
	}
}

func someRecursive(a []int, f callback) bool {
	return sRH(a, f, 0)
}

func sRH(a []int, f callback, i int) bool {
	n := len(a)
	if i == n - 1 {
		return f(a[i])
	} else {
		if f(a[i]) {
			return true
		} else {
			return sRH(a, f, i + 1)
		}
	}
}