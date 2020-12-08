package bubble

func Sort(slc *[]int) {
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
