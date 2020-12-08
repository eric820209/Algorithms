package main

import (
	"fmt"
	"mySort/selection"
)

func main() {
	var nums = []int{1, 5, 7, 89, 3, 7, 45}
	selection.Sort(&nums)
	fmt.Println(nums)
}
