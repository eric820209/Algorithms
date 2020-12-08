package selection

func Sort(slc *[]int) {
	var nums = *slc
	var length = len(nums)

	var m int
	for i := 0; i < length-1; i++ {
		m = i
		for s := i + 1; s < length; s++ {
			if nums[m] > nums[s] {
				m = s
			}
		}
		var temp = nums[m]
		for j := m; j > i; j-- {
			nums[j] = nums[j-1]
		}
		nums[i] = temp
	}
}
