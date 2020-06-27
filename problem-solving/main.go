package main

import(
	"fmt"
	"strings"
)

//explore examples -- simple, complex, empty inputs, invalid inputs
//break it down
func countCharacters(s string) map[string]int{
	m := make(map[string]int)
	n := len(s)
	for i := 0; i < n; i++ {
		if _, ok := m[strings.ToLower(string(s[i]))]; ok {
			m[strings.ToLower(string(s[i]))] += 1
		} else {
			m[strings.ToLower(string(s[i]))] = 1
		}
	}
	return m
}

func main() {
	var a = []int{-1,-2,-3,-4,-5}
	fmt.Println(maxSubArraySum(a,2))
}

// Patterns

/**
	* Frequency Counter
	*
	* Use maps to avoid nested loops
*/
func isAnagram(a,b string) bool {
	n := len(a)
	m := len(b)

	if m != n {
		return false
	}

	ia := true
	am := make(map[string]int)
	bm := make(map[string]int)

	for i := 0; i < n; i++ {
		if _, ok := am[strings.ToLower(string(a[i]))]; ok {
			am[strings.ToLower(string(a[i]))] += 1
		} else {
			am[strings.ToLower(string(a[i]))] = 1
		}
	}

	for i := 0; i < n; i++ {
		if _, ok := bm[strings.ToLower(string(b[i]))]; ok {
			bm[strings.ToLower(string(b[i]))] += 1
		} else {
			bm[strings.ToLower(string(b[i]))] = 1
		}
	}

	for key, val := range am {
		if bval, ok := bm[key]; ok {
			if val != bval {
				ia = false
			}
		} else {
			ia = false
		}
	}
	return ia
}

/**
	* Multiple pointers
	*
	* 
*/
//the array must be sorted for this to work
func sumZero(a []int) []int {
	l := 0
	r := len(a) - 1

	for l < r {
		v := a[l] + a[r] 
    if v == 0 {
			return []int{a[l],a[r]}
		} else if v < 0 {
			l++
		} else {
			r--
		}
	}

	return []int{}
}

func countUniqueValues(a []int) int{
	p1 := 1
	p2 := 0
	if len(a) == 0 {
		return 0
	}
	for p1 < len(a) {
		if a[p1] != a[p2] {
			p2++
			a[p2] = a[p1]
		}
		p1++
	}
	//this way, we can also easily return what the unique values are
	return p2+1
}

/**
	* Window Slider
	*
	* Used for a subset of data that exhibits a continuous property
*/
func maxSubArraySum(a []int, n int) int{
	if n > len(a) {
		return 0
	}
	l := 0
	r := 1
	cs := a[l]

	for r < n {
		cs += a[r]
		r++
	}
	max := cs
	for r < len(a) {
		cs = cs + a[r] - a[l]
		r++
		l++
		if cs > max {
			max = cs 
		}
	}

	return max
}