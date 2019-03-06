package main

import "fmt"

/*
	questions to ask:
	- only 0,1,2? yes
*/

/*
	1st approach: bucket sort

	Time	O(n)
	Space	O(3)
	0 ms, faster than 100.00%
*/
func sortColors(nums []int) {
	bucket := []int{0, 0, 0}
	for _, num := range nums {
		bucket[num]++
	}
	index := 0
	for i := 0; i < len(bucket); i++ {
		if bucket[i] > 0 {
			for j := 0; j < bucket[i]; j++ {
				nums[index] = i
				index++
			}
		}
	}
}

/*
	2nd approach: bucket sort but just dont init an array LOL

	Time	O(n)
	Space	O(3)
	0 ms, faster than 100.00%
*/
func sortColors1(nums []int) {
	zero := 0
	one := 0
	two := 0
	for _, num := range nums {
		if num == 0 {
			zero++
		} else if num == 1 {
			one++
		} else if num == 2 {
			two++
		}
	}
	index := 0
	if zero > 0 {
		for i := 0; i < zero; i++ {
			nums[index] = 0
			index++
		}
	}
	if one > 0 {
		for i := 0; i < one; i++ {
			nums[index] = 1
			index++
		}
	}
	if two > 0 {
		for i := 0; i < two; i++ {
			nums[index] = 2
			index++
		}
	}
}

/*
	2nd approach: merge sort

	Time	O(nlogn)
	Space	O(n)
	0 ms, faster than 100.00%
*/
func sortColors0(nums []int) {
	temp := mergeSort(nums)
	for i := 0; i < len(nums); i++ {
		nums[i] = temp[i]
	}
}

func mergeSort(nums []int) []int {
	// divide
	if len(nums) == 1 {
		return nums
	}
	mean := len(nums) / 2
	arr1 := mergeSort(nums[:mean])
	arr2 := mergeSort(nums[mean:])
	// merge 2 sorted arrays
	i, j := 0, 0
	result := []int{}
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	if i < len(arr1) {
		result = append(result, arr1[i:]...)
	}
	if j < len(arr2) {
		result = append(result, arr2[j:]...)
	}
	return result
}

func main() {
	a := []int{2, 0, 1, 2, 0, 1}
	sortColors1(a)
	fmt.Println(a)

	a = []int{}
	sortColors(a)
	fmt.Println(a)

	a = []int{1, 2, 1}
	sortColors(a)
	fmt.Println(a)
}
