package main

import "fmt"

func main(){
	fmt.Println(moveZeroes([]int{0,1,0,3,12}))
	fmt.Println(moveZeroes([]int{1,0}))
	fmt.Println(moveZeroes([]int{1}))
}

func moveZeroes(nums []int) []int{
	p1 := 0
	p2 := 0

	for p2 < len(nums) && p1 < len(nums) {
		if nums[p1] == 0 {
			if nums[p2] != 0 {
				nums[p1] = nums[p2]
				nums[p2] = 0
				p1++
			}
			p2++
		} else {
			p1++
			p2++
		}
	}
	return nums
}