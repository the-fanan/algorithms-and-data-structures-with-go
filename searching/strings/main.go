package main

import (
	"fmt"
)

func main() {
	fmt.Println(kmp("na", "fanan"))
}

func naiveSearch(n, h string) int {
	for i := 0; i < len(h); i++ {
		k := i
		for j := 0; j < len(n); j++ {
			if h[k] == n[j] {
				if j == len(n) - 1 {
					return i
				}
				k++
				continue
			} else {
				k++
				break
			}
		}
	}
	return -1
}

//character-jump heuristic
func boyerMoore(needle, haystack string) int {
	n := len(needle)
	h := len(haystack)
	lid := make(map[byte]int)
	//create a look up dictionary for last index of characters
	for i := 0; i < n; i++ {
		lid[needle[i]] = i
	}

	i,j := n - 1, n -1

	for i < h {
		if needle[j] == haystack[i] {
			//return if match found
			if j == 0 {
				return i
			}
			//traverse backwards
			i--
			j--
		} else {
			if val, ok := lid[haystack[i]]; ok {
				//match of character is seen ij needle at index val
				//i += n - min(val + 1, j)
				if (val + 1) < j {
					i += (n - (val + 1))
				} else {
					i += n - j
				}
				j = n - 1
			} else {
				//no match move to end of J 
				j = n - 1
				i += n
			}
		}
	}

	return -1
}


func kmp(needle, haystack string) int {
	n := len(needle)
	h := len(haystack)
	fail := computeKmpFail(needle)

	i := 0
	j := 0

	for i < h {
		if haystack[i] == needle[j] {
			if j == n - 1 {
				return i - n + 1
			}
			i++
			j++
		} else {
			if j > 0 {
				j = fail[j - 1]
			} else {
				i++
			}
		}
	}
	return -1
}

/**
* Returns longest prefix of s that is also a suffix
*/
//[f a n a n]
//[0 0 0 0 0]
//[p o l o i n p o l o o u t]
//[0 0 0 0 0 0 1 2 3 4 0 0 0]
//[a m a l g a m a t i o n]
//[0 0 1 0 0 1 2 3 0 0 0 0]
//[a m a l g a m r t i o n]
//[0 0 1 0 0 1 2 0 0 0 0 0]
func computeKmpFail(s string) []int {
	f := make([]int,len(s))
	i := 0
	j := 1

	for j < len(s) {
		if s[i] == s[j] {
			f[j] = f[j - 1] + 1
			i++
			j++
		} else if i > 0 {
			i = f[f[j - 1]]
		} else {
			j++
		}
	}
	return f
}