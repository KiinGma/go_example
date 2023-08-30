package waitGroup_example

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func A() {
	time.Sleep(time.Second * 5)
	fmt.Println("A执行完毕")
	wg.Done() //通知执行完毕
}

func B() {
	fmt.Println("B执行完毕")
	wg.Done() //通知执行完毕
}

var wg sync.WaitGroup

func TestWaitGroup(t *testing.T) {
	//需要等待协程执行完成的个数
	wg.Add(2)
	go A()
	go B()
	wg.Wait()
	fmt.Println("结束")
}
