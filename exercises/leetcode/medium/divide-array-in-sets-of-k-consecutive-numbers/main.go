package main

import "fmt"

func main(){
	nums := []int{3,2,1,2,3,4,3,4,5,9,10,11}
	fmt.Println(isPossibleDivideOptimized(nums,3))
}

/*
	* Given an array of integers nums and a positive integer k, find whether it's possible to divide this array into sets of k consecutive numbers
	* Return True if its possible otherwise return False.

	* Example 1:
	*	Input: nums = [1,2,3,3,4,4,5,6], k = 4
	* Output: true
	* Explanation: Array can be divided into [1,2,3,4] and [3,4,5,6].

	* Example 2:
	* Input: nums = [3,2,1,2,3,4,3,4,5,9,10,11], k = 3
	* Output: true
	* Explanation: Array can be divided into [1,2,3] , [2,3,4] , [3,4,5] and [9,10,11].

	* Example 3:
	* Input: nums = [3,3,2,2,1,1], k = 3
	* Output: true

	* Example 4:
	* Input: nums = [1,2,3,4], k = 3
	* Output: false
	* Explanation: Each array should be divided in subarrays of size 3.
*/
func isPossibleDivideOptimized(nums []int, k int) bool {
	min := int(^uint(0) >> 1)
	max := 0
	numbers := make(map[int]int)
	//convert to a map and find smallest number
	for _,v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		if _,ok := numbers[v]; !ok {
			numbers[v] = 1
		} else {
			numbers[v]++
		}
	}

	keepSearching := true 
	subL := 0
	running := min
	for keepSearching {
		if min > max {
			break
		}
		if subL == 0 {
			if v, ok := numbers[min]; ok {
				if v > 0 {
					subL++
					numbers[min]--
					running =  min + 1
					if numbers[min] == 0 {
						min++
					}
				} else {
					min++
				}
			} else {
				min++
			}
		} else {
			if v, ok := numbers[running]; ok {
				if v > 0 {
					subL++
					numbers[running]--
					if running == min && numbers[running] == 0{
						min++
					}
					running++

					if subL == k {
						subL = 0
					}
				} else {
					break
				}
			} else {
				//the number does not exist
				break
			}
		}
	}
	if subL > 0 {
		return false
	}
	return true
}

func isPossibleDivide(nums []int, k int) bool {
	//sort nums
	nums = mergeSort(nums)
	subL := 0
	currentInt := 0
	i := 0 
	j := 0
	n := len(nums)
	for j < n {
		if i >= n {
			return false
		}
		if subL == 0 {
			currentInt = nums[i]
			nums[i] = 0
			subL++
			i++
			j++
			continue
		}
		//follows the order
		if nums[i] == currentInt + 1 {
			currentInt = nums[i]
			nums[i] = 0
			subL++
			if i == j {
				j++
			}
			i++
		} else {
			if nums[i] == 0 {
				if i == j {
					j++
				}
				i++
			} else {
				i++
			}
		}

		if subL == k {
			subL = 0
			i = j
		}
	}
	if subL != 0 {
		return false
	}
	return true
}

func mergeSort(nums []int) []int {
	tempArr := make([]int, len(nums))
	recMergeSort(&nums,&tempArr,0,len(nums) - 1)
	return nums
}

func recMergeSort(nums *[]int,temp *[]int,l,r int){
	if l == r {
		return
	} else {
		m := (l + r) / 2
		recMergeSort(nums,temp,l,m)
		recMergeSort(nums,temp,m + 1,r)
		Merger(nums,temp,l,m + 1,r + 1)
	}
}

func Merger(nums *[]int,temp *[]int,l,r,end int){
	left := l 
	right := r 
	t := l 
	for left < r && right < end {
		if (*nums)[left] < (*nums)[right] {
			(*temp)[t] = (*nums)[left]
			left++
		} else {
			(*temp)[t] = (*nums)[right]
			right++
		}
		t++
	}

	for left < r {
		(*temp)[t] = (*nums)[left]
		left++
		t++
	}

	for right < end {
		(*temp)[t] = (*nums)[right]
		right++
		t++
	}

	for i := l; i < end; i++{
		(*nums)[i] = (*temp)[i]
	}
}