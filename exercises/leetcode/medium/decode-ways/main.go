package main

import "fmt"

func main(){
	fmt.Println(numDecodings("1787897759966261825913315262377298132516969578441236833255596967132573482281598412163216914566534565"))
}

/*
	* A message containing letters from A-Z is being encoded to numbers using the following mapping:

	* 'A' -> 1
	* 'B' -> 2
	* ...
	* 'Z' -> 26

	* Given a non-empty string containing only digits, determine the total number of ways to decode it.

	* Example 1:
	* Input: "12"
	* Output: 2
	* Explanation: It could be decoded as "AB" (1 2) or "L" (12).

	* Example 2:
	* Input: "226"
	* Output: 3
	* Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
*/

func numDecodings(s string) int {
	if s == "" {
		return 1
	}
	if string(s[0]) == "0" {
		return 0
	}
	m := make(map[string]int)
  return decodeHelper(s,&m)
}

func decodeHelper(s string, m *map[string]int) int {
	if string(s[0]) == "0" {
		return 0
	}

	if len(s) == 1 {
		if s == "0" {
			(*m)[s] = 0
			return 0
		} else {
			(*m)[s] = 1
			return 1
		}
	}
	if len(s) == 2 {
		if string(s[0]) == "0" {
			(*m)[s] = 0
			return 0
		}

		if convStI(string(s[0])) > 2 && string(s[1]) == "0" {
			(*m)[s] = 0
			return 0
		}

		if convStI(string(s[0])) > 2 && convStI(string(s[1])) > 0 {
			(*m)[s] = 1
			return 1
		}

		if string(s[0]) == "2" && convStI(string(s[1])) > 6 {
			(*m)[s] = 1
			return 1
		}

		if convStI(string(s[0])) <= 2 && string(s[1]) == "0" {
			(*m)[s] = 1
			return 1
		}
		(*m)[s] = 2
		return 2
	}

	ways := 0
	_,s1 := sliceString(s,1)
	s2s,s2 := sliceString(s,2)
	w1 := 0
	w2 := 0

	if v, ok := (*m)[s1]; ok {
		w1 = v
	} else {
		w1 = decodeHelper(s1,m)
	}

	if convStI(s2s) > 26 {
		w2 = 0
	} else {
		if v, ok := (*m)[s2]; ok {
			w2 = v
		} else {
			w2 = decodeHelper(s2,m)
		}
	}
	ways = w1 + w2
	return ways
}

func sliceString(s string, n int) (string,string) {
	sb := []byte(s)
	sbe := sb[n:]
	sbs := sb[:n]
	return string(sbs),string(sbe)
}

func convStI(s string) int {
	num := 0;
	for i := 0; i < len(s); i++ {
		n := int(s[i]) - 48 
		num += n 
		if i != len(s) - 1 {
			num *= 10
		}
	}
	return num
}