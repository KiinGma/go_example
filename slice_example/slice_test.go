package slice_example

import (
	"fmt"
	"testing"
)

// slice的扩容机制
func TestSliceDilatation(t *testing.T) {
	slice1 := []int{}
	oldCap := cap(slice1)
	for i := 0; i < 2048; i++ {
		slice1 = append(slice1, i)
		newCap := cap(slice1)
		if newCap != oldCap {
			fmt.Printf("%p %p cap=%d\n", &slice1, slice1, cap(slice1))
			oldCap = newCap
		}
	}
}
