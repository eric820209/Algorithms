package main

import (
	"Sort/bubble"
	"fmt"
)

func main() {
	var nums = []int{1, 5, 7, 89, 3, 7, 45}
	bubble.Sort(&nums)
	fmt.Println(nums)
}
