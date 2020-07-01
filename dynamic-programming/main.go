package main

import (
	_"fmt"
)

/**
	* Two condtions for Dynamic Programming
	*
	* 1. Overlapping sub problems -- smaller pieces that are repeated
	* 2. Optimal substructure -- the final solution depends on the optimal solution of it's smaller pices
*/
func main(){
	
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	m := make([]int, n + 1)
	m[1] = 1
	m[2] = 2
	return recCS(n, &m)
}

func recCS(n int, m *[]int) int{
	if n == 1 {
			return (*m)[1]
	}else if n == 2{
			return (*m)[2]
	} else {
			l := 0
			r := 0
			
			if l = (*m)[n - 1]; (*m)[n - 1] == 0 {
					l = recCS(n - 1, m)
			}
			
			if r = (*m)[n - 2]; (*m)[n - 2] == 0 {
					l = recCS(n - 2, m)
			}
			(*m)[n] = l + r
			return l + r
	}
}

func sumSubArr(a []int, i,j int) int{
	s := 0
	for i < j {
		s += a[i]
		i++
	}
	return s
}

//solving Fibonacci by memoization (top-down approach) -- start from top and go down to base case
func dynamicFib(n int) int {
	m := make([]int,n)
	r := fibHelper(n, &m)
	return r
}

func fibHelper(n int, m *[]int) int{
	if n == 1{
		(*m)[n - 1] = 1
		return 1
	} else if n == 2 {
		(*m)[n - 1] = 1
		return 1
	} else {
		l := 0
		r := 0
		//the order of finding the nearest sub-problem is very important for an optimal solution
		if l = (*m)[n - 2]; (*m)[n - 2] == 0 {
			l = fibHelper(n - 1, m)
		}

		if r = (*m)[n - 3]; (*m)[n - 3] == 0 {
			r = fibHelper(n - 2, m)
		}

		(*m)[n - 1] = l + r
		return (*m)[n - 1]
	}
}
//tabulation method (bottom-up) -- start from base case and go up to optimal solution 
//it improves space complexity because call stack is not loaded with multiple recursive functions
func fibTab(n int) int {
	tab := make([]int, n + 1)
	tab[1] = 1
	tab[2] = 1

	for i := 3; i <= n; i++ {
		tab[i] = tab[i-1] + tab[i-2]
	}
	return tab[n]
}

//other functions
func longestCommonSubsequence(x,y string) [][]int {
	xn := len(x)
	yn := len(y)

	l := make([][]int,(xn + 1))
	for i:= 0; i <= xn; i++ {
		l[i] = make([]int,(yn + 1))
	}
	//no matter what happens, we must build a table then traverse
	for i := 0; i < xn; i++{
		for j := 0; j < yn; j++{
			if x[i] == y[j] {
				l[i+1][j+1] = l[i][j] + 1
			} else {
				//take the max
				if l[i+1][j+1] = l[i][j+1]; l[i+1][j] > l[i][j+1]  {
					l[i+1][j+1] = l[i+1][j]
				}
			}
		}
	}
	return l
}

func subseqSol(x,y string, l [][]int) string{
	s := make([]byte,0)
	i := len(x)
	j := len(y)
	for l[i][j] > 0 {
		if x[i-1] == y[j-1] {
			s = append(s,x[i-1])
			i--
			j--
		} else if l[i-1][j] >= l[i][j-1] {
			i--
		} else {
			j--
		}
	} 

	return string(ReverseStringC([]byte(s)))
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