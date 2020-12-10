package main

import (
	"fmt"
	"mySort/merge"
)

func main() {
	var nums = []int{1, 5, 7, 89, 3, 7, 45}
	//insertion.Sort(&nums)
	nums = merge.Sort(nums)
	fmt.Println(nums)
}
