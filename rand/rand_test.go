package rand

import (
	"fmt"
	"testing"
)

func TestInt64n(t *testing.T) {
	n, err := Int64n(500)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func TestFloat64n(t *testing.T) {
	n, err := Float64()
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
