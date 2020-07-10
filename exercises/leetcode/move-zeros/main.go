package main

import "fmt"

func main(){
	fmt.Println(moveZeroes([]int{0,1,0,3,12}))
	fmt.Println(moveZeroes([]int{1,0}))
	fmt.Println(moveZeroes([]int{1}))
}

func moveZeroes(nums []int) []int{
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j];
			i++
		}
	}
	for k := i; k < len(nums); k++{
		nums[k] = 0
	}
	return nums
}