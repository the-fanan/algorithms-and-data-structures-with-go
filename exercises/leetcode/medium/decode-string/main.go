package main

import (
	"fmt"
)

func main(){
	fmt.Println(decodeString("10[ff2[a]]2[e]fs"))
}
const NUM = "NUM"
const ALPHA = "APLHA"

func decodeString(s string) string {
	decoded := ""
	stringList := make([]string,0)
	mult := 0
	bracketStack := make([]string,0)
	needsDecoding := false 
	decodingString := ""
	runningString := ""
	runningType := ""

	for i := 0; i < len(s); i++ {
		if mult == 0 {

			if runningString == "" {
				runningString = string(s[i])
				if isAlpha(s[i]) {
					runningType = ALPHA
				} else {
					runningType = NUM
				}
			} else {
				if isAlpha(s[i]) && runningType == NUM {
					//a switch from number to string
					mult = convStI(runningString)
					runningString = string(s[i])
					runningType = ALPHA
				} else if isAlpha(s[i]) && runningType == ALPHA {
					//continuation from string to string
					runningString += string(s[i])
				} else if isNum(s[i]) && runningType == ALPHA {
					//switch from string to number
					stringList = append(stringList,runningString)
					runningString = string(s[i])
					runningType = NUM
				} else if isNum(s[i]) && runningType == NUM {
					//continuation from number to number
					runningString += string(s[i])
				} else if string(s[i]) == "[" {
					bracketStack = append(bracketStack,"[")
					decodingString += string(s[i])
					mult = convStI(runningString)
					runningString = ""
					runningType = ""
				}
			}
		} else {
			if string(s[i]) == "[" {
				bracketStack = append(bracketStack,"[")
				decodingString += string(s[i])
			} else if string(s[i]) == "]" {
				lastBracket := bracketStack[len(bracketStack) - 1]
				bracketStack = bracketStack[:len(bracketStack) - 1]
				if len(bracketStack) == 0 && lastBracket == "[" {
					//we have completed the decoded string
					dsb := []byte(decodingString)
					dsb = dsb[1:]
					decodingString = string(dsb)
					if needsDecoding {
						stringList = append(stringList, repeatString(decodeString(decodingString), mult))
					} else {
						stringList = append(stringList, repeatString(decodingString, mult))
					}
					decodingString = ""
					mult = 0
					needsDecoding = false
				} else {
					needsDecoding = true 
					decodingString += string(s[i])
				}
			} else {
				decodingString += string(s[i])
			}
		}
	}

	if runningString != "" {
		stringList = append(stringList, runningString)
	}

	for _,v := range stringList {
		decoded += v
	}

	return decoded
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
func isNum(s byte) bool {
	if int(s) >= 48 && int(s) <= 57{
		return true
	}
	return false
}

func isAlpha(s byte) bool {
	if int(s) >= 97 && int(s) <= 122{
		return true
	}
	return false
}

func repeatString(s string, n int) string {
	fs := s
	mGI := n/2
	mSL := len(s) * mGI

	for len(fs) <= mSL {
		fs = fs + fs
	}

	fsL := len(fs)
	sL := len(s)
	eFSL := sL * n

	if (fsL < eFSL) {	
		newN := (eFSL/sL) - (fsL/sL)
		fs += repeatString(s, newN)
	}
	return fs
}
