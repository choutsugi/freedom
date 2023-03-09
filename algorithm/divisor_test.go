package algorithm

import (
	"fmt"
	"testing"
)

func TestMaxDivisor4Slice(t *testing.T) {
	nums := []int64{6, 12, 18}
	divisor := MaxDivisor4Slice(nums)
	fmt.Println(divisor)
}
