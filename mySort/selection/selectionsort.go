package selection

func Sort(slc *[]int) {
	var nums = *slc
	var length = len(nums)

	var m int
	for i := 0; i < length-1; i++ {
		m = i
		var temp = nums[i]
		for s := i + 1; s < length; s++ {
			if nums[m] > nums[s] {
				m = s
			}
		}
		nums[i] = nums[m]
		nums[m] = temp
	}
}
