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
	fmt.Println(countCharacters("Fanan123"))
}