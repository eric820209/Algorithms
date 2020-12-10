package merge

func Sort(nums []int) []int {

	var length = len(nums)
	if length == 1 {
		return nums
	}

	var mid = length / 2
	if length%2 == 1 {
		mid += 1
	}
	var left = nums[0:mid]
	var right = nums[mid:]

	return merge(Sort(left), Sort(right))
}

func merge(left, right []int) []int {
	var mixlen = len(left) + len(right)
	var mix = make([]int, mixlen)
	var ileft, iright int
	for ileft < len(left) && iright < len(right) {
		if left[ileft] <= right[iright] {
			mix[ileft+iright] = left[ileft]
			ileft++
		} else {
			mix[ileft+iright] = right[iright]
			iright++
		}
	}
	if ileft == len(left) {
		mix = append(mix[0:ileft+iright], right[iright:]...)
		return mix
	} else {
		mix = append(mix[0:ileft+iright], left[ileft:]...)
		return mix
	}
}
