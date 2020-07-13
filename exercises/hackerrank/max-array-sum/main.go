package main

import "fmt"

func main(){
		fmt.Println(maxSubsetSum([]int32{3,5,-7,8,10}))
}

func maxSubsetSum(arr []int32) int32 {
	dp := make([]int32,len(arr))
	dp[0] = arr[0]
	max := arr[0]
	dp[1] = arr[1]
	if dp[1] > max {
		max = dp[1]
	}
	for i := 2; i < len(arr); i++ {
			if arr[i] > arr[i] + dp[i-2] {
				dp[i] = arr[i]
			} else {
				dp[i] = arr[i] + dp[i-2]
			}

			if dp[i] > max {
				max = dp[i]
			} else {
				dp[i] = max
			}
	}
	return int32(max)
}
