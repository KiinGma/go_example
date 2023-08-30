package chan_example

import (
	"fmt"
	"testing"
	"time"
)

// 基础 chan 的建立
func TestChanCreate(t *testing.T) {
	c := make(chan int, 6)
	c <- 1
	fmt.Println(<-c)
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
	fmt.Println(<-c)
}

// chan 写 panic 情况
// 当写操作执行时,要么存在缓存,要么在别的协程要先有读堵塞
func TestChanPanicWithReErr1(t *testing.T) {
	c := make(chan int, 0)
	go func() {
		fmt.Println(<-c)
	}()
	c <- 1
	time.Sleep(time.Second * 5)
}

// chan 写 堵塞时关闭 panic 情况
// 以下情况为 关闭通道后还对 chan 写入 panic
func TestChanPanicWithReErr2(t *testing.T) {
	c := make(chan int, 6)
	go func() {
		c <- 1
	}()

	go func() {
		c <- 2
	}()

	go func() {
		c <- 3
	}()
	close(c)
}

// (未复现)关闭channel时会把recvq中的G全部唤醒，本该写入G的数据位置为nil。把sendq中的G全部唤醒，但这些G会panic。
func TestChanPanicWithReErr3(t *testing.T) {
	c := make(chan int)
	go func() {
		fmt.Println("1 : ", <-c)
	}()

	go func() {
		c <- 1
	}()
	go func() {
		c <- 2
	}()

	go func() {
		time.Sleep(time.Second * 2)
		close(c)
	}()

	go func() {
		fmt.Println("2 : ", <-c)
	}()

	time.Sleep(5 * time.Second)

}

// 关闭已经关闭的chan
func TestChanPanicWithCloseErr1(t *testing.T) {
	c := make(chan int)
	close(c)
	close(c)
}

// 关闭没有初始化的chan
func TestChanPanicWithCloseErr2(t *testing.T) {
	var c chan int
	close(c)
}
