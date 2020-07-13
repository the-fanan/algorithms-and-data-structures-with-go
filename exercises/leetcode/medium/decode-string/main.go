package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(repeatString("f", 5))
}
const MULT = "MULT"
const ADD = "ADD"
const NUM = "NUM"
const ALPHA = "ALPHA"

func decodeString(s string) string {
	decoded := ""
	multStack := make([]int,0)
	stringStack := make([]string,0)
	oppStack := make([]string,0)//ADD,MULT
	running := ""
	runningType := ""//NUM,ALPH
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "[" {
			//multiply
			n,_:= strconv.Atoi(running)
			multStack = append(multStack,n)
			oppStack = append(oppStack, MULT)
			running = ""
			runningType = ""
		} else if string(s[i]) == "]" {
			//add
			if running != "" {
				//guard against repeated closing tags
				stringStack = append(stringStack, running)
				oppStack = append(oppStack, ADD)
				running = ""
				runningType = ""
			}
		} else {
			if running == "" {
				running = string(s[i])
				if isNum(s[i]) {
					runningType = NUM
				} else {
					runningType = ALPHA
				}
			} else {
				if (runningType == NUM && isNum(s[i])) ||  (runningType == ALPHA && isAlpha(s[i])){
					running += string(s[i])
				} else {
					//at this point only number or letter and mismatch will appear
					if isNum(s[i]) {
						//we have encountered a transition from letters to numbers
						stringStack = append(stringStack, running)
						oppStack = append(oppStack, ADD)
						running = string(s[i])
						runningType = NUM
					} else {
						//we have encountered a transition from numbers to letters
						n,_:= strconv.Atoi(running)
						multStack = append(multStack,n)
						//oppStack = append(oppStack, MULT)
						running = string(s[i])
						runningType = ALPHA
					}
				}
			}
		}
		oppStack = oppStack[:len(oppStack) - 1]//take out the last add opp 
		for len(multStack) > 0 {
			opp := oppStack[len(oppStack) - 1]
			oppStack = oppStack[:len(oppStack) - 1]
			if opp == ADD {
				str1 := stringStack[len(stringStack) - 1]
				stringStack = stringStack[:len(stringStack) - 1]
				str2 := stringStack[len(stringStack) - 1]
				stringStack = stringStack[:len(stringStack) - 1]
				str2 += str1 
				stringStack = append(stringStack,str2)
			} else {
				//opp will me mult

			}
		}
	}
	return decoded
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
