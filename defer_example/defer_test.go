package defer_example

import (
	"fmt"
	"testing"
)

// 先注册先出
func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(2)
	}()
	defer func() {
		fmt.Println(3)
	}()
}

//和return的执行顺序 , 先defer 后 Return

func TestDeferAndReturn(t *testing.T) {
	fmt.Println(re())
}

func re() int {
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(2)
	}()
	defer func() {
		fmt.Println(3)
	}()
	return 4
}

// defer能否能改变return的值? 能
// 这种特性只适用于具名返回值，对于没有具名返回值的匿名返回值，defer语句无法改变其值。
func TestDeferChangeReturnValue(t *testing.T) {
	i := DeferChangeReturnValue()
	fmt.Println(i)
}

func DeferChangeReturnValue() (val int) {
	defer func() {
		val = 2
	}()
	return 1
}

// 1.panic之后会执行defer
// 2.先出的defer panic 也不会阻止后输出的defer
// 3.panic后不会再执行任何函数
func TestPanicAndDefer(t *testing.T) {
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		panic(nil)
	}()
	fmt.Println(2)
	panic(nil)
	fmt.Println(3)
	return
}
