package insertion

func Sort(slc *[]int) {
	var nums = *slc
	var length = len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j > 0; j-- {
			if nums[j] < nums[j-1] {
				var temp = nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = temp
			}
		}
	}
}
