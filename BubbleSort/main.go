package main

import "fmt"

func main() {
	var nums = []int{1, 5, 7, 89, 3, 7, 45}
	BubbleSort(&nums)
	fmt.Println(nums)
}

func BubbleSort(slc *[]int) {
	var nums = *slc
	var length = len(nums)

	for t := 0; t < length; t++ {
		for i := 0; i < length-t-1; i++ {
			if nums[i] > nums[i+1] {
				var temp = nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = temp
			}
		}
	}
}
