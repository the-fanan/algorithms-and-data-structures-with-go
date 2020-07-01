package main

import (
	"fmt"
)

func main(){
//	fmt.Println(LongestCommonSubsequence("bsbininm", "jmjkbkjkv"))
//	fmt.Println(MaximumSubarray([]int{-1,2,-3,-4}))
	fmt.Println(MaximumProductSubarray([]int{-2,3,-4}))
}

func LongestCommonSubsequence(a,b string) int{
	x := a
	y := b
	xn := len(x)
	yn := len(y)

	l := make([][]int,(xn + 1))
	for i:= 0; i <= xn; i++ {
		l[i] = make([]int,(yn + 1))
	}
	for i := 0; i < xn; i++{
		for j := 0; j < yn; j++{
			if x[i] == y[j] {
				l[i+1][j+1] = l[i][j] + 1
			} else {
				if l[i+1][j+1] = l[i][j+1]; l[i][j+1] < l[i+1][j] {
					l[i+1][j+1] = l[i+1][j]
				}
			}
		}
	}
	return l[xn][yn]
}

func MaximumSubarray(a []int) int{
	maxValue := 0
	runningMax := 0
	startI := 0

	for i := 0; i < len(a); i++{
		if i == 0 {
			maxValue = a[i]
			runningMax = a[i]
		} else {
			//give me the max between number alone or number plus running max
			if runningMax + a[i] > a[i] {
				runningMax = a[i] + runningMax
			} else {
				runningMax = a[i]
				startI = i
			}
		}
		if runningMax > maxValue {
			maxValue = runningMax
		}
	}

	fmt.Println(startI)
	return maxValue
}

func MaximumProductSubarray(nums []int) int{
	n := len(nums)
	max := 0
	for i := 0; i < n; i++{
		localMax := 1
		for j := i; j < n; j++{
			if j == 0 {
				max = nums[j]
			}
			localMax *= nums[j]
			if localMax > max {
				max = localMax
			}
		}
	}
		
	return max
}