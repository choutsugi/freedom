package sugar

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	res := Contains(nums, 5)
	fmt.Println(res)
}

func TestDelete(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	res := Delete(nums, 5)
	fmt.Println(res)
}

func TestDiff(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{1, 2, 5}
	res := Diff(nums1, nums2)
	fmt.Println(res)
}
