package sugar

import (
	"fmt"
	"testing"
)

func TestTernary(t *testing.T) {

	numA := 1
	numB := 2

	res := Ternary(numA == numB, 1, 3)
	fmt.Println(res)
}
