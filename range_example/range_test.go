package range_example

import (
	"fmt"
	"testing"
	"time"
)

// 使用range读取chan值
func TestRangeReadeChan(t *testing.T) {
	c := make(chan int, 1)
	go func() {
		c <- 3
	}()
	go chanRange(c)
	time.Sleep(time.Minute)
}

func chanRange(chanName chan int) {
	for e := range chanName {
		fmt.Printf("Get element from chan: %d\n", e)
	}
}
