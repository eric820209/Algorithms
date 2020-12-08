package bubble

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var nums = []int{1, 5, 8, 2, 3}
	Sort(&nums)
	var exp = []int{1, 2, 3, 5, 8}
	var same = reflect.DeepEqual(nums, exp)
	if !same {
		t.Error("Not same")
	}
}
