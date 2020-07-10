package main

import "fmt"

func main(){
	fmt.Println(intersection([]int{1,2,2,2,4,5,6},[]int{1,4,4,6,7,9,6}))
}

func intersection(nums1 []int, nums2 []int) []int {
	intersection := make(map[int]int)
	for _,v := range nums1 {
			if _, ok := intersection[v]; !ok {
					intersection[v] = 1
			}
	}
	
	for _, v := range nums2 {
			if _,ok := intersection[v]; ok {
					intersection[v]++
			}
	}
	
	result := make([]int,0)
	for number,count := range intersection {
			if count > 1 {
					result = append(result,number)
			}
	}
	return result
}