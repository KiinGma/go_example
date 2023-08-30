package select_example

import (
	"fmt"
	"testing"
	"time"
)

func TestSelectCreate2(t *testing.T) {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c3 := make(chan int, 1)

	go func() {
		c1 <- 1
	}()
	go func() {
		c2 <- 2
	}()
	go func() {
		c3 <- 3
	}()
	go func() {
		for {
			select {
			case i := <-c1:
				fmt.Println(i)
			case i := <-c2:
				fmt.Println(i)
			case i := <-c3:
				fmt.Println(i)
			}
		}
	}()
	time.Sleep(time.Minute)
}
